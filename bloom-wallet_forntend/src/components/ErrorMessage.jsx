import React from 'react';
import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';

export default function ErrorMessage({ title = '请求失败', message, onRetry }) {
  return (
    <Alert
      severity="error"
      variant="outlined"
      sx={{
        borderColor: 'rgba(255,255,255,0.12)',
        backgroundColor: 'rgba(255,255,255,0.03)',
      }}
    >
      <AlertTitle>{title}</AlertTitle>
      <Stack spacing={1} direction={{ xs: 'column', sm: 'row' }} alignItems={{ sm: 'center' }}>
        <span>{message || '请稍后重试。'}</span>
        {onRetry ? (
          <Button size="small" color="inherit" variant="outlined" onClick={onRetry}>
            重试
          </Button>
        ) : null}
      </Stack>
    </Alert>
  );
}
