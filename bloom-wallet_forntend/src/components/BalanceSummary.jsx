import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';

export default function BalanceSummary({ currentAccount }) {
  return (
    <Card>
      <CardContent>
        <Stack spacing={1.25}>
          <Typography variant="overline" sx={{ color: 'rgba(255,255,255,0.65)' }}>
            钱包概览
          </Typography>
          <Stack
            direction={{ xs: 'column', sm: 'row' }}
            spacing={2}
            alignItems={{ sm: 'center' }}
            justifyContent="space-between"
          >
            <Stack spacing={0.5}>
              <Typography variant="h5" sx={{ fontWeight: 900, letterSpacing: 0.2 }}>
                总资产
              </Typography>
              <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)' }}>
                当前后端未提供余额接口，先展示账户信息（后续可扩展链上余额）。
              </Typography>
            </Stack>
            <Typography
              variant="h4"
              sx={{
                fontWeight: 900,
                background: 'linear-gradient(135deg, #7c4dff, #00e5ff)',
                WebkitBackgroundClip: 'text',
                WebkitTextFillColor: 'transparent',
              }}
            >
              N/A
            </Typography>
          </Stack>

          <Divider sx={{ borderColor: 'rgba(255,255,255,0.08)' }} />

          <Stack direction="row" spacing={1} justifyContent="space-between">
            <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)' }}>
              当前账户
            </Typography>
            <Typography variant="body2" sx={{ fontWeight: 700 }}>
              {currentAccount?.name || (currentAccount ? `Account #${currentAccount.index}` : '—')}
            </Typography>
          </Stack>
        </Stack>
      </CardContent>
    </Card>
  );
}
