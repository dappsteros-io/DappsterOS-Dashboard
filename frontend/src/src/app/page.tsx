"use client";

import NoLayout from "@/components/NoLayout";
import { Card, Col, Row, Statistic } from "antd";
import { LikeOutlined } from "@ant-design/icons";
import React from "react";
import { Typography } from "antd"; 

const Index: React.FC = () => {
  return (
    <NoLayout>
      <div className="w-full flex flex-col justify-center text-2xl mt-40 gap-4 h-96 text-center p-4">
        <Typography.Title>Welcome to DappsterOS Dashboard!</Typography.Title>
        <Row gutter={16} className="justify-center">
          <Col span={8}>
            <Card bordered={false}>
              <Statistic
                title="Feedback"
                value={1128}
                prefix={<LikeOutlined />}
              />
            </Card>
          </Col>
          <Col span={8}>
            <Card bordered={false}>
              <Statistic title="Stable" value={99.8} suffix="/ 100" />
            </Card>
          </Col>
        </Row>
      </div>
    </NoLayout>
  );
};

export default Index;
