import React, { useMemo, useState } from 'react';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Stack from '@mui/material/Stack';
import Tab from '@mui/material/Tab';
import Tabs from '@mui/material/Tabs';

import MainLayout from './layouts/MainLayout.jsx';
import DashboardPage from './pages/DashboardPage.jsx';
import AccountsPage from './pages/AccountsPage.jsx';
import ReceivePage from './pages/ReceivePage.jsx';
import WalletSetupPage from './pages/WalletSetupPage.jsx';

const routes = [
  { key: 'dashboard', label: '概览' },
  { key: 'accounts', label: '账户' },
  { key: 'receive', label: '收款' },
  { key: 'setup', label: '创建/恢复' },
];

export default function App() {
  const [route, setRoute] = useState('dashboard');

  const page = useMemo(() => {
    if (route === 'accounts') return <AccountsPage />;
    if (route === 'receive') return <ReceivePage />;
    if (route === 'setup') return <WalletSetupPage />;
    return <DashboardPage />;
  }, [route]);

  return (
    <MainLayout>
      <Container maxWidth="lg" sx={{ py: { xs: 2, md: 4 } }}>
        <Stack spacing={2.5}>
          <Box
            sx={{
              position: 'sticky',
              top: 72,
              zIndex: 1,
              backdropFilter: 'blur(10px)',
              backgroundColor: 'rgba(11,15,26,0.55)',
              border: '1px solid rgba(255,255,255,0.06)',
              borderRadius: 3,
              px: 1,
            }}
          >
            <Tabs
              value={route}
              onChange={(_, v) => setRoute(v)}
              variant="scrollable"
              scrollButtons="auto"
            >
              {routes.map((r) => (
                <Tab key={r.key} value={r.key} label={r.label} />
              ))}
            </Tabs>
          </Box>

          {page}
        </Stack>
      </Container>
    </MainLayout>
  );
}
