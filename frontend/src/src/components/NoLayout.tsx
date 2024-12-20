"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";

import {
  ConfigProvider,
  Layout,
  Menu,
  Avatar,
  Button,
  theme,
  Flex,
  Dropdown,
  Affix,
  Drawer,
} from "antd";
import type { MenuProps } from "antd";
import { WebSocketProvider } from "next-ws/client";
import { MenuInfo } from "@/types";

import {
  MoonOutlined,
  SunOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  SettingOutlined,
  UserOutlined,
  LinkedinOutlined,
  LogoutOutlined,
  UserAddOutlined,
} from "@ant-design/icons";

import {
  BsRobot,
  BsSpeedometer2,
  BsShield,
  BsPersonLinesFill,
  BsNewspaper,
  BsCheck2Square,
  BsPeopleFill,
} from "react-icons/bs";
import Image from "next/image";
import Logo from "../assets/img/logo.png";
import LayoutContext from "@/contexts/LayoutContextProvider";
import { wsURL } from "@/api/index";
import { ROUTE } from "@/routes";

const { Header, Sider } = Layout;

const sideBarItems: MenuProps["items"] = [
  {
    key: "home",
    icon: <BsRobot />,
    label: `Home`,
  },
  {
    key: "dashboard",
    icon: <BsSpeedometer2 />,
    label: `Dashboard`,
  },
];

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter();

  const profileItems: MenuProps["items"] = [
    {
      key: "myprofile",
      icon: <UserOutlined />,
      label: `Profile`,
    },
    {
      key: "setting",
      icon: <SettingOutlined />,
      label: `Setting`,
    },
    {
      key: "logout",
      icon: <LogoutOutlined />,
      label: `Logout`,
      onClick: () => {
        console.log("logout");
        router.push("/auth/signin");
      },
    },
  ];

  const {
    token: { colorBgContainer },
  } = theme.useToken();

  const [collapsed, setCollapsed] = useState<boolean>(false);
  // const [selectedkeys, setSelectedkeys] = useState<string[]>(["home"]);
  const [isDarkTheme, setDarkTheme] = useState<boolean>(true);

  const onThemeChange = () => {
    if (!isDarkTheme) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
    setDarkTheme(!isDarkTheme);
  };

  return (
    <ConfigProvider
      theme={{
        token: {
          borderRadius: 0,
          motion: false,
        },
        algorithm: isDarkTheme ? theme.darkAlgorithm : theme.defaultAlgorithm,
      }}
    >
      <LayoutContext.Provider value={{ collapsed, setCollapsed }}>
        <WebSocketProvider url={wsURL}>
          <Layout className="min-h-screen">
            <Header
              className="p-0 flex justify-between items-center px-4"
              style={{ background: isDarkTheme ? "" : colorBgContainer }}
            >
              <div
                className="h-16 justify-center items-center text-center cursor-pointer flex flex-row gap-4"
                onClick={() => {
                  // setSelectedkeys(["home"]);
                  router.push(ROUTE.landing);
                }}
              >
                <Image src={Logo} alt="L" width={48} height={48} priority />
                <h2 className={collapsed ? `hidden` : ""}>DappsterOS</h2>
              </div>
              <Button
                type="text"
                icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                onClick={() => setCollapsed(!collapsed)}
                className="!w-16 h-16 text-base invisible"
              />
              <Flex
                align="flex-end"
                justify="space-between"
                className="gap-2 items-center"
              >
                <Button
                  type="text"
                  icon={isDarkTheme ? <SunOutlined /> : <MoonOutlined />}
                  onClick={onThemeChange}
                  className="!w-16 h-16 text-base"
                />
                <Button
                  type="link"
                  icon={<UserOutlined />}
                  onClick={() => router.push(ROUTE.signin)}
                  className=" text-base"
                >
                  Login
                </Button>
                <Button
                  type="link"
                  icon={<UserAddOutlined />}
                  onClick={() => router.push(ROUTE.signup)}
                  className=" text-base"
                >
                  Register
                </Button>
              </Flex>
            </Header>
            {children}
          </Layout>
        </WebSocketProvider>
      </LayoutContext.Provider>
    </ConfigProvider>
  );
}
