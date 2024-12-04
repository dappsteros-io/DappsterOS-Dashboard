"use client";

import React, { useEffect, useState } from "react";
import type { FormProps } from "antd";
import {
  Button,
  Checkbox,
  Form,
  Input,
  Card,
  Layout,
  Flex,
} from "antd";
import { useRouter } from "next/navigation";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { signIn } from "@/store/user/actions";
import { ROUTE } from "@/routes";
import Link from "next/link";
import Loader from "@/components/Loader"; 
import { MdOutlineAlternateEmail, MdPassword } from "react-icons/md";

type FieldType = {
  email?: string;
  password?: string;
  remember?: string;
};

const SignIn: React.FC = () => {
  const router = useRouter();
  const dispatch = useAppDispatch();
  const [error, setError] = useState<Error>();
  const err = useAppSelector((state) => state.user.error);
  const onFinish: FormProps<FieldType>["onFinish"] = (values) => {
    dispatch(signIn(values)).then((res) => {
      if (res.meta.requestStatus == "fulfilled") {
        router.push(ROUTE.dashboard);
      }
    });
  };

  const onFinishFailed: FormProps<FieldType>["onFinishFailed"] = (
    errorInfo
  ) => {
    console.log("Failed:", errorInfo);
  };
  const [loading, setLoading] = useState<boolean>(true);
  useEffect(() => {
    setLoading(false);
  }, []);
  useEffect(() => {
    setError(err);
  }, [err]);

  return loading ? (
    <Loader />
  ) : (
    <Layout className="flex items-center justify-center">
      <Card style={{ width: 360 }} title="Sign In">
        <Form
          name="signin"
          style={{ maxWidth: 600 }}
          initialValues={{ remember: true }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
        >
          <Form.Item<FieldType>
            name="email"
            rules={[{ required: true, message: "Please input your email!" }]}
            validateStatus={error?.message ? "error" : ""}
          >
            <Input
              type="email"
              autoComplete="off"
              prefix={<MdOutlineAlternateEmail />}
            />
          </Form.Item>

          <Form.Item<FieldType>
            name="password"
            rules={[{ required: true, message: "Please input your password!" }]}
            validateStatus={error?.message ? "error" : ""}
          >
            <Input.Password prefix={<MdPassword />} />
          </Form.Item>

          <Form.Item label={null}>
            <Flex justify="space-between" align="center">
              <Form.Item name="remember" valuePropName="checked" noStyle>
                <Checkbox>Remember me</Checkbox>
              </Form.Item>
              <Button type="link" href="">
                Forgot password
              </Button>
            </Flex>
          </Form.Item>

          <Form.Item label={null}>
            <Button type="primary" htmlType="submit" className="w-full">
              Submit
            </Button>
          </Form.Item>
          <Form.Item label={null}>
            or{" "}
            <Link href="/auth/signup">
              <Button type="link">Register now!</Button>
            </Link>
          </Form.Item>
        </Form>
      </Card>
    </Layout>
  );
};

export default SignIn;
