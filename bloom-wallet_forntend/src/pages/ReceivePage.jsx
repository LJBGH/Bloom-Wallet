import React, { useCallback, useEffect, useState } from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';

import AddressDisplay from '../components/AddressDisplay.jsx';
import Loader from '../components/Loader.jsx';
import ErrorMessage from '../components/ErrorMessage.jsx';
import { getCurrentAccount } from '../services/walletService';

export default function ReceivePage() {
  const [loading, setLoading] = useState(true);
  const [err, setErr] = useState('');
  const [current, setCurrent] = useState(null);

  const load = useCallback(async () => {
    setLoading(true);
    setErr('');
    try {
      const cur = await getCurrentAccount();
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

  if (loading) return <Loader label="加载收款信息…" />;
  if (err) return <ErrorMessage message={err} onRetry={load} />;

  return (
    <Stack spacing={2.5}>
      <Card>
        <CardContent>
          <Stack spacing={1.25}>
            <Typography variant="h6" sx={{ fontWeight: 900 }}>
              收款
            </Typography>
            <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)' }}>
              将该地址提供给对方即可完成转入。后续可以在此页增加二维码、链选择等。
            </Typography>
            <Divider sx={{ borderColor: 'rgba(255,255,255,0.08)' }} />
            <AddressDisplay address={current?.address} showShort={false} />
          </Stack>
        </CardContent>
      </Card>
    </Stack>
  );
}
