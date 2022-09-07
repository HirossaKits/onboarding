import React from 'react';
import { Flex, Icon, Link, FlexProps } from '@chakra-ui/react';
import { IconType } from 'react-icons';
import { ReactText } from 'react';

interface NavMenuProps extends FlexProps {
  icon: IconType;
  children: ReactText;
}
const NavMenu = ({ icon, children, ...rest }: NavMenuProps) => {
  return (
    <Link
      href="#"
      style={{ textDecoration: 'none' }}
      _focus={{ boxShadow: 'none' }}
    >
      <Flex
        align="center"
        p="4"
        mx="4"
        borderRadius="lg"
        role="group"
        cursor="pointer"
        _hover={{
          bg: 'cyan.400',
          color: 'white',
        }}
        {...rest}
      >
        {icon && (
          <Icon
            mr="4"
            fontSize="16"
            _groupHover={{
              color: 'white',
            }}
            as={icon}
          />
        )}
        {children}
      </Flex>
    </Link>
  );
};

export default NavMenu;
