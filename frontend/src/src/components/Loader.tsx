import { Flex, Spin, SpinProps } from "antd";
import { LoadingOutlined } from "@ant-design/icons";

export default function Loader(props: SpinProps) {
  return (
    <Flex className="fixed top-0 left-0 bg-opacity-20 bg-gray-600 w-screen h-screen z-50">
      <Spin
        className="fixed top-1/2 left-1/2 z-50 -translate-x-1/2 -translate-y-1/2"
        indicator={<LoadingOutlined style={{ fontSize: 48 }} />}
      ></Spin>
    </Flex>
  );
}
