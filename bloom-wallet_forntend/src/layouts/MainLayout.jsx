import React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import ContentCopyRoundedIcon from '@mui/icons-material/ContentCopyRounded';

import { getBackendBaseUrl } from '../config/env';

export default function MainLayout({ children }) {
  const backend = getBackendBaseUrl();

  async function copy(text) {
    try {
      await navigator.clipboard.writeText(text);
    } catch {
      // ignore
    }
  }

  return (
    <Box sx={{ minHeight: '100%', background: 'radial-gradient(1200px 600px at 20% 0%, rgba(124,77,255,0.20), rgba(11,15,26,0))' }}>
      <AppBar
        position="sticky"
        elevation={0}
        sx={{
          borderBottom: '1px solid rgba(255,255,255,0.08)',
          backgroundColor: 'rgba(11,15,26,0.72)',
          backdropFilter: 'blur(10px)',
        }}
      >
        <Toolbar sx={{ minHeight: 72 }}>
          <Stack direction="row" alignItems="center" spacing={1} sx={{ flex: 1 }}>
            <Box
              sx={{
                width: 28,
                height: 28,
                borderRadius: 2,
                background: 'linear-gradient(135deg, #7c4dff, #00e5ff)',
              }}
            />
            <Typography variant="h6" sx={{ fontWeight: 700, letterSpacing: 0.2 }}>
              Bloom Wallet
            </Typography>
          </Stack>

          <Stack direction="row" alignItems="center" spacing={0.5}>
            <Typography
              variant="caption"
              sx={{ color: 'rgba(255,255,255,0.70)', display: { xs: 'none', sm: 'block' } }}
            >
              API: {backend}
            </Typography>
            <Tooltip title="复制 API 地址">
              <span>
                <IconButton size="small" onClick={() => copy(backend)}>
                  <ContentCopyRoundedIcon fontSize="inherit" />
                </IconButton>
              </span>
            </Tooltip>
          </Stack>
        </Toolbar>
      </AppBar>

      {children}
    </Box>
  );
}
