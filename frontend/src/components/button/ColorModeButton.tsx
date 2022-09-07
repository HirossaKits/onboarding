import React from 'react';
import { IconButton, useColorMode } from '@chakra-ui/react';
import { FiMoon, FiSun } from 'react-icons/fi';

const ColorModeButton = () => {
  const { colorMode, toggleColorMode } = useColorMode();
  return (
    <IconButton
      size="lg"
      variant="ghost"
      aria-label="open menu"
      icon={colorMode === 'light' ? <FiMoon /> : <FiSun />}
      onClick={toggleColorMode}
    />
  );
};

export default ColorModeButton;
