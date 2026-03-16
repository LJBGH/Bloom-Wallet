import React, { useCallback, useEffect, useState } from 'react';
import Grid from '@mui/material/Grid';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Divider from '@mui/material/Divider';
import Box from '@mui/material/Box';

import BalanceSummary from '../components/BalanceSummary.jsx';
import AccountCard from '../components/AccountCard.jsx';
import Loader from '../components/Loader.jsx';
import ErrorMessage from '../components/ErrorMessage.jsx';
import { getAllAccounts, getCurrentAccount, switchCurrentAccount } from '../services/walletService';
import { getAuthToken, setAuthToken } from '../services/apiClient';

export default function DashboardPage() {
  const [loading, setLoading] = useState(true);
  const [err, setErr] = useState('');
  const [accounts, setAccounts] = useState([]);
  const [current, setCurrent] = useState(null);

  const [tokenInput, setTokenInput] = useState(getAuthToken());

  const load = useCallback(async () => {
    setLoading(true);
    setErr('');
    try {
      const [all, cur] = await Promise.all([getAllAccounts(), getCurrentAccount()]);
      setAccounts(all);
      setCurrent(cur);
    } catch (e) {
      setErr(e?.message || String(e));
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    load();
  }, [load]);

  async function onSwitch(index) {
    try {
      await switchCurrentAccount(index);
      await load();
    } catch (e) {
      setErr(e?.message || String(e));
    }
  }

  function saveToken() {
    setAuthToken(tokenInput.trim());
    load();
  }

  return (
    <Stack spacing={2.5}>
      <BalanceSummary currentAccount={current} />

      <Card>
        <CardContent>
          <Stack spacing={1.5}>
            <Typography variant="subtitle1" sx={{ fontWeight: 800 }}>
              认证 Token
            </Typography>
            <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)' }}>
              后端接口需要 `Authorization: Bearer token`。你可以先在 Swagger 登录获取 token，然后粘贴到这里。
            </Typography>
            <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
              <TextField
                fullWidth
                label="Bearer Token（可直接粘贴 token 或 Bearer xxx）"
                value={tokenInput}
                onChange={(e) => setTokenInput(e.target.value)}
              />
              <Button variant="contained" onClick={saveToken} sx={{ minWidth: 120 }}>
                保存并刷新
              </Button>
            </Stack>
            <Divider sx={{ borderColor: 'rgba(255,255,255,0.08)' }} />
            <Typography variant="caption" sx={{ color: 'rgba(255,255,255,0.55)' }}>
              提示：该 token 保存在浏览器 localStorage（key: bloom_wallet_token）。
            </Typography>
          </Stack>
        </CardContent>
      </Card>

      {loading ? <Loader label="加载账户信息…" /> : null}
      {err ? <ErrorMessage message={err} onRetry={load} /> : null}

      {!loading && !err ? (
        <Stack spacing={1.25}>
          <Typography variant="h6" sx={{ fontWeight: 900 }}>
            账户列表
          </Typography>
          {accounts.length === 0 ? (
            <Box
              sx={{
                border: '1px dashed rgba(255,255,255,0.18)',
                borderRadius: 3,
                py: 4,
                px: 3,
                textAlign: 'center',
                color: 'rgba(255,255,255,0.70)',
              }}
            >
              暂无账户数据。你可以先调用后端的创建/导入/恢复钱包接口，然后刷新本页。
            </Box>
          ) : (
            <Grid container spacing={2}>
              {accounts.map((a) => (
                <Grid key={a.index ?? a.address} item xs={12} sm={6} md={4}>
                  <AccountCard
                    account={a}
                    isCurrent={current?.index === a.index}
                    onSwitch={onSwitch}
                  />
                </Grid>
              ))}
            </Grid>
          )}
        </Stack>
      ) : null}
    </Stack>
  );
}
