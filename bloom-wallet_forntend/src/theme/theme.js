import { createTheme } from '@mui/material/styles';

const base = {
  shape: { borderRadius: 16 },
  typography: {
    fontFamily: ['Roboto', 'system-ui', 'Segoe UI', 'Arial', 'sans-serif'].join(','),
  },
  components: {
    MuiCard: {
      defaultProps: {
        variant: 'outlined',
      },
      styleOverrides: {
        root: {
          borderColor: 'rgba(255,255,255,0.08)',
          backgroundImage: 'none',
        },
      },
    },
  },
};

export const darkTheme = createTheme({
  ...base,
  palette: {
    mode: 'dark',
    primary: { main: '#7c4dff' },
    secondary: { main: '#00e5ff' },
    background: {
      default: '#0b0f1a',
      paper: 'rgba(255,255,255,0.04)',
    },
  },
});
