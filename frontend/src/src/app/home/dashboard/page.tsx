"use client";

import { ClusterData } from "@/types/cluster";
import type { GetProp, TableProps, TreeDataNode, GetProps } from "antd";
import { Badge, Button, Flex, Space, Table, Tree } from "antd";
import React, { useEffect, useState } from "react";

import { DownOutlined } from "@ant-design/icons";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { getNodes, getVMs, startVM, stopVM } from "@/store/proxmox/actions";
import { NodeData, VMData } from "@/types/proxmox";
import {
  BsDatabase,
  BsPlay,
  BsPlayFill,
  BsRepeat,
  BsStop,
} from "react-icons/bs";
import { RiRefreshLine } from "react-icons/ri";

import Loader from "@/components/Loader";
type DirectoryTreeProps = GetProps<typeof Tree.DirectoryTree>;

const { DirectoryTree } = Tree;

type SizeType = TableProps["size"];
type ColumnsType<T extends object> = GetProp<TableProps<T>, "columns">;
type TablePagination<T extends object> = NonNullable<
  Exclude<TableProps<T>["pagination"], boolean>
>;
type TablePaginationPosition = NonNullable<
  TablePagination<any>["position"]
>[number];
type ExpandableConfig<T extends object> = TableProps<T>["expandable"];
type TableRowSelection<T extends object> = TableProps<T>["rowSelection"];

const defaultExpandable: ExpandableConfig<ClusterData> = {
  expandedRowRender: (record: ClusterData) => <p>{record.description}</p>,
};

const defaultTitle = () => "Here is title";
const defaultFooter = () => "Here is footer";

const Index: React.FC = () => {
  const [bordered, setBordered] = useState(false);
  const [loading, setLoading] = useState(false);
  const [size, setSize] = useState<SizeType>("large");
  const [expandable, setExpandable] =
    useState<ExpandableConfig<ClusterData>>(defaultExpandable);
  const [showTitle, setShowTitle] = useState(false);
  const [showHeader, setShowHeader] = useState(true);
  const [showFooter, setShowFooter] = useState(true);
  const [rowSelection, setRowSelection] = useState<
    TableRowSelection<ClusterData> | undefined
  >({});
  const [hasData, setHasData] = useState(true);
  const [tableLayout, setTableLayout] = useState();
  const [ellipsis, setEllipsis] = useState(false);
  const [yScroll, setYScroll] = useState(false);
  const [xScroll, setXScroll] = useState<string>();
  const nodes = useAppSelector((state) => state.proxmox.nodes);
  const vms = useAppSelector((state) => state.proxmox.vms);
  const dispatch = useAppDispatch();

  const columns: ColumnsType<VMData> = [
    {
      title: "VMID",
      dataIndex: "VMID",
      sorter: (a, b) => a.VMID - b.VMID,
    },
    {
      title: "Name",
      dataIndex: "Name",
      sorter: true,
    },
    {
      title: "Node",
      dataIndex: "Node",
    },
    {
      title: "Tags",
      dataIndex: "tags",
      render: (value) =>
        value
          ?.split(",")
          .map((v: string, k: number) => <Badge count={v} key={k} />),
    },
    {
      title: "Status",
      dataIndex: "Status",
      sorter: true,
    },
    {
      title: "Action",
      key: "action",
      render: (val, rec, index) => (
        <Space size="middle">
          <Button
            icon={
              rec.Status == "stopped" ? (
                <BsPlayFill color="green" />
              ) : (
                <BsStop color="red" />
              )
            }
            onClick={() =>
              rec.Status == "stopped" ? onStartVM(rec.VMID) : onStopVM(rec.VMID)
            }
          />
          {/* <a>
            <Space>
              <DownOutlined />
            </Space>
          </a> */}
        </Space>
      ),
    },
  ];
  const tableColumns = columns.map((item) => ({ ...item, ellipsis }));

  if (xScroll === "fixed") {
    tableColumns[0].fixed = true;
    tableColumns[tableColumns.length - 1].fixed = "right";
  }

  const scroll: { x?: number | string; y?: number | string } = {};
  if (yScroll) {
    scroll.y = 240;
  }
  if (xScroll) {
    scroll.x = "100vw";
  }
  const tableProps: TableProps<ClusterData> = {
    bordered,
    loading,
    size,
    expandable,
    title: showTitle ? defaultTitle : undefined,
    showHeader,
    footer: showFooter ? defaultFooter : undefined,
    rowSelection,
    scroll,
    tableLayout,
  };
  const onStartVM = (vmid: number) => {
    dispatch(startVM(vmid)).then((res) => {
      setLoading(true);
      dispatch(getVMs({})).then((res) => {
        setLoading(false);
      });
    });
  };
  const onStopVM = (vmid: number) => {
    dispatch(stopVM(vmid)).then((res) => {
      setLoading(true);
      dispatch(getVMs({})).then((res) => {
        setLoading(false);
      });
    });
  };
  const onSelect: DirectoryTreeProps["onSelect"] = (keys, info) => {
    console.log("Trigger Select", keys, info);
    setLoading(true);
    dispatch(getVMs({ node: info.node.key })).then((res) => {
      console.log({ res });
      setLoading(false);
    });
  };

  const onExpand: DirectoryTreeProps["onExpand"] = (keys, info) => {
    console.log("Trigger Expand", keys, info);
  };
  const onRefresh = () => {
    setLoading(true);
    dispatch(getVMs({})).then((res) => {
      setLoading(false);
    });
  };
  useEffect(() => {
    setLoading(true);
    dispatch(getNodes({})).then((res) => {
      setLoading(false);
    });
    setLoading(true);
    dispatch(getVMs({})).then((res) => {
      setLoading(false);
    });
  }, []);

  return (
    <div className="w-full text-lg text-center p-4">
      <Flex flex={1} gap={4}>
        {/*         <Flex flex={1} title="Nodes" className="flex-col">
          <h3 className="block">Nodes</h3>
          <div>
            <DirectoryTree
              showIcon={false}
              // showLine
              // multiple
              className="p-4 min-w-48"
              // draggable
              titleRender={(node) => (
                <div className="flex flex-row">
                  <BsDatabase className="w-4 h-4 my-auto mr-2" />
                  {node.title}
                </div>
              )}
              defaultExpandAll
              onSelect={onSelect}
              onExpand={onExpand}
              treeData={nodes.map((n) => ({
                title: n.Node,
                key: n.ID,
              }))}
            />
          </div>
        </Flex> */}
        <Flex align="start" className="flex-col">
          <Space style={{ marginBottom: 16 }}>
            <h3>Machines</h3>
            <div className="w-full"></div>
            <Button icon={<RiRefreshLine />} onClick={onRefresh}></Button>
          </Space>
          <Table<VMData>
            pagination={{ position: ["none", "bottomCenter"] }}
            columns={tableColumns}
            dataSource={vms}
            scroll={{ x: "max-content", y: 100 * 5 }}
            loading={loading}
          />
        </Flex>
      </Flex>
    </div>
  );
};

export default Index;