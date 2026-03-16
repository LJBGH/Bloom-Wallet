import React, { useCallback, useEffect, useMemo, useState } from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Divider from '@mui/material/Divider';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

import Loader from '../components/Loader.jsx';
import ErrorMessage from '../components/ErrorMessage.jsx';
import AddressDisplay from '../components/AddressDisplay.jsx';
import { getAllAccounts, getCurrentAccount, switchCurrentAccount } from '../services/walletService';

export default function AccountsPage() {
  const [loading, setLoading] = useState(true);
  const [err, setErr] = useState('');
  const [accounts, setAccounts] = useState([]);
  const [current, setCurrent] = useState(null);

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

  const currentDetail = useMemo(() => current, [current]);

  async function onSwitch(index) {
    try {
      await switchCurrentAccount(index);
      await load();
    } catch (e) {
      setErr(e?.message || String(e));
    }
  }

  if (loading) return <Loader label="加载账户列表…" />;
  if (err) return <ErrorMessage message={err} onRetry={load} />;

  return (
    <Stack spacing={2.5}>
      <Card>
        <CardContent>
          <Stack spacing={1}>
            <Typography variant="h6" sx={{ fontWeight: 900 }}>
              当前账户
            </Typography>
            <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)' }}>
              当前账户用于“收款地址”等功能的默认选择。
            </Typography>
            <Divider sx={{ borderColor: 'rgba(255,255,255,0.08)' }} />
            <AddressDisplay address={currentDetail?.address} showShort={false} />
            <Typography variant="caption" sx={{ color: 'rgba(255,255,255,0.55)' }}>
              Index: {currentDetail?.index ?? '—'}；Path: {currentDetail?.path || '—'}
            </Typography>
          </Stack>
        </CardContent>
      </Card>

      <Card>
        <CardContent>
          <Stack spacing={1.5}>
            <Typography variant="h6" sx={{ fontWeight: 900 }}>
              所有账户
            </Typography>
            <List disablePadding>
              {accounts.map((a, idx) => {
                const isCur = current?.index === a.index;
                return (
                  <React.Fragment key={a.index ?? a.address ?? idx}>
                    <ListItem
                      disableGutters
                      secondaryAction={
                        <Button
                          size="small"
                          variant={isCur ? 'contained' : 'outlined'}
                          disabled={isCur}
                          onClick={() => onSwitch(a.index)}
                        >
                          {isCur ? '当前' : '切换'}
                        </Button>
                      }
                    >
                      <ListItemText
                        primary={a.name || `Account #${a.index}`}
                        secondary={a.address}
                        primaryTypographyProps={{ sx: { fontWeight: 800 } }}
                        secondaryTypographyProps={{
                          sx: { color: 'rgba(255,255,255,0.65)' },
                        }}
                      />
                    </ListItem>
                    {idx < accounts.length - 1 ? (
                      <Divider sx={{ borderColor: 'rgba(255,255,255,0.08)' }} />
                    ) : null}
                  </React.Fragment>
                );
              })}
            </List>
          </Stack>
        </CardContent>
      </Card>
    </Stack>
  );
}
