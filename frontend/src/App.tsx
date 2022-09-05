import React from 'react';
import { ChakraProvider } from '@chakra-ui/react';

import Navbar from './components/navbar/Navbar';

export default function App() {
  return (
    <ChakraProvider>
      <Navbar />
    </ChakraProvider>
  );
}
