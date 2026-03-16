import React from 'react';
import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';

export default function Loader({ label = '加载中…', minHeight = 120 }) {
  return (
    <Box sx={{ minHeight, display: 'grid', placeItems: 'center' }}>
      <Stack spacing={1} alignItems="center">
        <CircularProgress size={28} />
        <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.72)' }}>
          {label}
        </Typography>
      </Stack>
    </Box>
  );
}
