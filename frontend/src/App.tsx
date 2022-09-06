import React from 'react';
import { ChakraProvider } from '@chakra-ui/react';

import Navbar from './components/navbar/Navbar';
import Sidebar from './components/sidebar/Sidebar';

export default function App() {
  return (
    <ChakraProvider>
      <Sidebar>
        <Navbar />
      </Sidebar>
    </ChakraProvider>
  );
}
