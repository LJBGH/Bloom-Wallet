import React, { useMemo, useState } from 'react';
import Box from '@mui/material/Box';
import IconButton from '@mui/material/IconButton';
import Snackbar from '@mui/material/Snackbar';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Tooltip from '@mui/material/Tooltip';
import ContentCopyRoundedIcon from '@mui/icons-material/ContentCopyRounded';

function shorten(addr) {
  if (!addr) return '';
  if (addr.length <= 14) return addr;
  return `${addr.slice(0, 8)}…${addr.slice(-6)}`;
}

export default function AddressDisplay({ address, label = '收款地址', showShort = true }) {
  const [open, setOpen] = useState(false);
  const display = useMemo(() => (showShort ? shorten(address) : address || ''), [address, showShort]);

  async function copy() {
    if (!address) return;
    try {
      await navigator.clipboard.writeText(address);
      setOpen(true);
    } catch {
      // ignore
    }
  }

  return (
    <Stack spacing={1}>
      <Box sx={{ display: 'flex', gap: 1, alignItems: 'center' }}>
        <TextField
          fullWidth
          label={label}
          value={display}
          InputProps={{ readOnly: true }}
          size="medium"
        />
        <Tooltip title="复制完整地址">
          <span>
            <IconButton onClick={copy} disabled={!address}>
              <ContentCopyRoundedIcon />
            </IconButton>
          </span>
        </Tooltip>
      </Box>

      <Snackbar
        open={open}
        autoHideDuration={1800}
        onClose={() => setOpen(false)}
        message="已复制地址"
      />
    </Stack>
  );
}
