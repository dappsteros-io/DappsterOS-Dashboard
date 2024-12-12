"use client";

import React from "react";
import type { FormProps } from "antd";
import { Button, Form, Input, Card, Layout } from "antd";
import { useRouter } from "next/navigation";
import { useAppDispatch } from "@/lib/hooks";
import { signIn, signUp } from "@/store/user/actions";
import Link from "next/link";

type FieldType = {
  email?: string;
  password?: string;
  password2?: string;
  remember?: string;
};

const SignUp: React.FC = () => {
  const router = useRouter();
  const dispatch = useAppDispatch();
  const onFinish: FormProps<FieldType>["onFinish"] = (values) => {
    dispatch(signUp(values));
    // router.push("/home")
  };

  const onFinishFailed: FormProps<FieldType>["onFinishFailed"] = (
    errorInfo
  ) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <Layout className="flex items-center justify-center">
      <Card style={{ width: 360 }} title="Sign Up">
        <Form
          name="signup"
          style={{ maxWidth: 600 }}
          initialValues={{ remember: true }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
        >
          <Form.Item<FieldType>
            name="email"
            rules={[{ required: true, message: "Please input your email!" }]}
          >
            <Input type="email" />
          </Form.Item>

          <Form.Item<FieldType>
            name="password"
            rules={[{ required: true, message: "Please input your password!" }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item<FieldType>
            name="password2"
            rules={[{ required: true, message: "Please input your password!" }]}
          >
            <Input.Password />
          </Form.Item>

          {/* <Form.Item label={null}>
            <Flex justify="space-between" align="center">
              <Form.Item name="remember" valuePropName="checked" noStyle>
                <Checkbox>Remember me</Checkbox>
              </Form.Item>
              <Button type="link" href="">Forgot password</Button>
            </Flex>
          </Form.Item> */}

          <Form.Item label={null}>
            <Button type="primary" htmlType="submit" className="w-full">
              Submit
            </Button>
          </Form.Item>
          <Form.Item label={null}>
            Already have an account?
            <Link href="/auth/signin">
              <Button type="link">
                Sign in
              </Button>
            </Link>
          </Form.Item>
        </Form>
      </Card>
    </Layout>
  );
};

export default SignUp;
