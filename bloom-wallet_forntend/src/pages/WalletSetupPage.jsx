import React, { useCallback, useMemo, useState } from 'react';
import Alert from '@mui/material/Alert';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Divider from '@mui/material/Divider';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

import ErrorMessage from '../components/ErrorMessage.jsx';
import Loader from '../components/Loader.jsx';
import AddressDisplay from '../components/AddressDisplay.jsx';
import {
  createMnemonic,
  createWallet,
  restoreWallet,
  restoreWithMnemonic,
  restoreWithPrivateKey,
  getAllAccounts,
  getCurrentAccount,
} from '../services/walletService';

function toText(x) {
  if (x == null) return '';
  if (typeof x === 'string') return x;
  try {
    return JSON.stringify(x, null, 2);
  } catch {
    return String(x);
  }
}

export default function WalletSetupPage() {
  const [busy, setBusy] = useState(false);
  const [err, setErr] = useState('');
  const [lastResult, setLastResult] = useState(null);

  const [generatedMnemonic, setGeneratedMnemonic] = useState('');
  const [createMnemonicInput, setCreateMnemonicInput] = useState('');
  const [restoreMnemonicInput, setRestoreMnemonicInput] = useState('');
  const [restorePrivateKeyInput, setRestorePrivateKeyInput] = useState('');

  const [accountsPreview, setAccountsPreview] = useState({ current: null, accounts: [] });

  const refreshPreview = useCallback(async () => {
    const [accounts, current] = await Promise.all([getAllAccounts(), getCurrentAccount()]);
    setAccountsPreview({ accounts, current });
  }, []);

  const run = useCallback(
    async (fn) => {
      setBusy(true);
      setErr('');
      try {
        const res = await fn();
        setLastResult(res);
        await refreshPreview();
        return res;
      } catch (e) {
        setErr(e?.message || String(e));
        throw e;
      } finally {
        setBusy(false);
      }
    },
    [refreshPreview],
  );

  const currentAddress = useMemo(() => accountsPreview.current?.address || '', [accountsPreview.current]);

  return (
    <Stack spacing={2.5}>
      <Card>
        <CardContent>
          <Stack spacing={1}>
            <Typography variant="h6" sx={{ fontWeight: 900 }}>
              创建与恢复钱包
            </Typography>
            <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)' }}>
              这些操作会调用后端 `/api/wallet/*` 接口。成功后会自动刷新“当前账户”和“账户列表”预览。
            </Typography>
            <Divider sx={{ borderColor: 'rgba(255,255,255,0.08)' }} />
            <Alert
              severity="info"
              variant="outlined"
              sx={{ borderColor: 'rgba(255,255,255,0.12)', backgroundColor: 'rgba(255,255,255,0.03)' }}
            >
              如果请求 401/403，请先到“概览”页设置 `Bearer token`（后端 Swagger 登录可获取）。
            </Alert>
          </Stack>
        </CardContent>
      </Card>

      {busy ? <Loader label="处理中…" /> : null}
      {err ? <ErrorMessage message={err} /> : null}

      <Stack direction={{ xs: 'column', md: 'row' }} spacing={2}>
        <Card sx={{ flex: 1 }}>
          <CardContent>
            <Stack spacing={1.5}>
              <Typography variant="subtitle1" sx={{ fontWeight: 900 }}>
                1) 生成助记词
              </Typography>
              <Button
                variant="contained"
                disabled={busy}
                onClick={() =>
                  run(async () => {
                    const res = await createMnemonic();
                    const m = res?.mnemonic || '';
                    setGeneratedMnemonic(m);
                    if (m) setCreateMnemonicInput(m);
                    return res;
                  })
                }
              >
                生成助记词
              </Button>
              <TextField
                label="生成的助记词"
                value={generatedMnemonic}
                fullWidth
                multiline
                minRows={2}
                InputProps={{ readOnly: true }}
              />
            </Stack>
          </CardContent>
        </Card>

        <Card sx={{ flex: 1 }}>
          <CardContent>
            <Stack spacing={1.5}>
              <Typography variant="subtitle1" sx={{ fontWeight: 900 }}>
                2) 创建钱包（可选助记词）
              </Typography>
              <TextField
                label="助记词（不填则后端自动生成）"
                value={createMnemonicInput}
                onChange={(e) => setCreateMnemonicInput(e.target.value)}
                fullWidth
                multiline
                minRows={2}
                placeholder="word1 word2 ... word12"
              />
              <Button
                variant="contained"
                disabled={busy}
                onClick={() =>
                  run(async () => {
                    return await createWallet({ mnemonic: createMnemonicInput });
                  })
                }
              >
                创建钱包
              </Button>
            </Stack>
          </CardContent>
        </Card>
      </Stack>

      <Card>
        <CardContent>
          <Stack spacing={2}>
            <Typography variant="subtitle1" sx={{ fontWeight: 900 }}>
              3) 恢复钱包
            </Typography>

            <Stack direction={{ xs: 'column', md: 'row' }} spacing={2}>
              <Box sx={{ flex: 1 }}>
                <TextField
                  label="通过助记词恢复"
                  value={restoreMnemonicInput}
                  onChange={(e) => setRestoreMnemonicInput(e.target.value)}
                  fullWidth
                  multiline
                  minRows={2}
                  placeholder="word1 word2 ... word12"
                />
                <Button
                  sx={{ mt: 1 }}
                  variant="outlined"
                  disabled={busy || !restoreMnemonicInput.trim()}
                  onClick={() => run(() => restoreWithMnemonic(restoreMnemonicInput))}
                >
                  按助记词恢复
                </Button>
              </Box>

              <Box sx={{ flex: 1 }}>
                <TextField
                  label="通过私钥恢复（hex 字符串）"
                  value={restorePrivateKeyInput}
                  onChange={(e) => setRestorePrivateKeyInput(e.target.value)}
                  fullWidth
                  placeholder="0x..."
                />
                <Button
                  sx={{ mt: 1 }}
                  variant="outlined"
                  disabled={busy || !restorePrivateKeyInput.trim()}
                  onClick={() => run(() => restoreWithPrivateKey(restorePrivateKeyInput))}
                >
                  按私钥恢复
                </Button>
              </Box>

              <Box sx={{ flex: 0.6 }}>
                <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)', mb: 1 }}>
                  使用后端默认备份源恢复
                </Typography>
                <Button variant="outlined" disabled={busy} onClick={() => run(() => restoreWallet())}>
                  默认恢复
                </Button>
              </Box>
            </Stack>
          </Stack>
        </CardContent>
      </Card>

      <Stack direction={{ xs: 'column', md: 'row' }} spacing={2}>
        <Card sx={{ flex: 1 }}>
          <CardContent>
            <Stack spacing={1.25}>
              <Typography variant="subtitle1" sx={{ fontWeight: 900 }}>
                刷新后的账户预览
              </Typography>
              <AddressDisplay address={currentAddress} showShort={false} label="当前账户地址" />
              <Typography variant="caption" sx={{ color: 'rgba(255,255,255,0.55)' }}>
                账户数量：{accountsPreview.accounts?.length ?? 0}
              </Typography>
              <Button variant="text" disabled={busy} onClick={() => run(() => refreshPreview())}>
                手动刷新预览
              </Button>
            </Stack>
          </CardContent>
        </Card>

        <Card sx={{ flex: 1 }}>
          <CardContent>
            <Stack spacing={1.25}>
              <Typography variant="subtitle1" sx={{ fontWeight: 900 }}>
                最近一次接口返回
              </Typography>
              <TextField
                value={lastResult ? toText(lastResult) : ''}
                placeholder="执行任意操作后会显示返回内容"
                multiline
                minRows={8}
                fullWidth
                InputProps={{ readOnly: true, sx: { fontFamily: 'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace' } }}
              />
            </Stack>
          </CardContent>
        </Card>
      </Stack>
    </Stack>
  );
}

