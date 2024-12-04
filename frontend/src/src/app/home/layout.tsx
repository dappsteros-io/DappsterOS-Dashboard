"use client";
import type { Metadata } from "next";
import localFont from "next/font/local";
import "../globals.css";
import Layout from "@/components/Layout";
import NoLayout from "@/components/NoLayout";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { useRouter } from "next/navigation";
import { ROUTE } from "@/routes";
import { useEffect, useState } from "react";
import { Flex, Spin } from "antd";
import { LoadingOutlined } from "@ant-design/icons";
import { getMe } from "@/store/user/actions";

export default function HomeLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [loading, setLoading] = useState<boolean>(true);
  const router = useRouter();
  const user = useAppSelector((state) => state.user);
  const dispatch = useAppDispatch();
  useEffect(() => {
    if (!user?.data?.id) {
      dispatch(getMe({})).then((res) => {
        if (res.meta.requestStatus !== "fulfilled") {
          // router.push(ROUTE.dashboard);
          router.push(ROUTE.signin);
        }
      });
    }
    setLoading(false);
  }, []);
  return loading ? (
    <Flex align="center" gap="middle">
      <Spin indicator={<LoadingOutlined spin />} size="large" />
    </Flex>
  ) : (
    <Layout>{children}</Layout>
  );
}
