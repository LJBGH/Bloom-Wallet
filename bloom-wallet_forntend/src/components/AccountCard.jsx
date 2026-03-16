import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Chip from '@mui/material/Chip';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Divider from '@mui/material/Divider';

function shorten(addr) {
  if (!addr) return '';
  if (addr.length <= 14) return addr;
  return `${addr.slice(0, 8)}…${addr.slice(-6)}`;
}

export default function AccountCard({ account, isCurrent, onSwitch }) {
  return (
    <Card
      sx={{
        height: '100%',
        borderColor: isCurrent ? 'rgba(124,77,255,0.55)' : 'rgba(255,255,255,0.08)',
        boxShadow: isCurrent ? '0 0 0 1px rgba(124,77,255,0.35)' : 'none',
      }}
    >
      <CardContent>
        <Stack spacing={1.5}>
          <Stack direction="row" spacing={1} alignItems="center" justifyContent="space-between">
            <Stack spacing={0.25}>
              <Typography variant="subtitle1" sx={{ fontWeight: 700 }}>
                {account?.name || `Account #${account?.index ?? '-'}`}
              </Typography>
              <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.70)' }}>
                {shorten(account?.address)}
              </Typography>
            </Stack>
            {isCurrent ? <Chip size="small" color="primary" label="当前" /> : null}
          </Stack>

          <Divider sx={{ borderColor: 'rgba(255,255,255,0.08)' }} />

          <Stack direction="row" spacing={1} alignItems="center" justifyContent="space-between">
            <Stack spacing={0.25}>
              <Typography variant="caption" sx={{ color: 'rgba(255,255,255,0.60)' }}>
                余额
              </Typography>
              <Typography variant="h6" sx={{ fontWeight: 800, letterSpacing: 0.2 }}>
                N/A
              </Typography>
            </Stack>
            <Button
              variant={isCurrent ? 'contained' : 'outlined'}
              color="primary"
              size="small"
              onClick={() => onSwitch?.(account?.index)}
              disabled={isCurrent || account?.index == null}
            >
              设为当前
            </Button>
          </Stack>
        </Stack>
      </CardContent>
    </Card>
  );
}
