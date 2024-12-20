"use client";

import { ClusterData } from "@/types/cluster";
import {
  GetProp,
  TableProps,
  TreeDataNode,
  GetProps,
  PopconfirmProps,
  Tooltip,
  Badge,
  Button,
  Flex,
  message,
  Popconfirm,
  Space,
  Table,
  Tree,
} from "antd";
import React, { useEffect, useState } from "react";

import {
  DownOutlined,
  PlusCircleFilled,
  PlusSquareOutlined,
} from "@ant-design/icons";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import {
  getNodes,
  getVMs,
  startVM,
  stopVM,
  deleteVM,
  createVM,
  installDappster,
  checkDappster,
} from "@/store/proxmox/actions";
import { NodeData, VMData } from "@/types/proxmox";
import {
  BsBackpack,
  BsCheck2All,
  BsDatabase,
  BsPlay,
  BsPlayFill,
  BsRepeat,
  BsStop,
  BsTrash,
} from "react-icons/bs";
import { RiRefreshLine } from "react-icons/ri";
import { VscVm } from "react-icons/vsc";

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
          <Tooltip title={rec.Status == "running" ? "Stop" : "Start"}>
            <Button
              icon={
                rec.Status == "running" ? (
                  <BsStop color="red" />
                ) : (
                  <BsPlayFill color="green" />
                )
              }
              onClick={() =>
                rec.Status == "running"
                  ? onStopVM(rec.VMID)
                  : onStartVM(rec.VMID)
              }
            />
          </Tooltip>
          <Tooltip title="Install DappsterOS">
            <Button
              disabled={rec.Status !== "running"}
              icon={<BsBackpack />}
              onClick={() => onInstallDappsterOS(rec.VMID)}
            />
          </Tooltip>
          <Tooltip title="Check DappsterOS">
            <Button
              disabled={rec.Status !== "running"}
              icon={<BsCheck2All />}
              onClick={() => onCheckDappsterOS(rec.VMID)}
            />
          </Tooltip>
          <Popconfirm
            title="Delete the VM"
            description="Are you sure to delete this VM?"
            onConfirm={() => onDeleteVM(rec.VMID)}
            onCancel={cancel}
            okText="Yes"
            cancelText="No"
          >
            <Tooltip title="Remove VM">
              <Button icon={<BsTrash color="red" />} />
            </Tooltip>
          </Popconfirm>
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
      onRefresh();
    });
  };

  const onStopVM = (vmid: number) => {
    dispatch(stopVM(vmid)).then((res) => {
      setLoading(true);
      onRefresh();
    });
  };

  const onDeleteVM = (vmid: number) => {
    dispatch(deleteVM(vmid)).then((res) => {
      setLoading(true);
      onRefresh();
    });
  };

  const onInstallDappsterOS = (vmid: number) => {
    setLoading(true);
    dispatch(installDappster(vmid)).then((res) => {
      setLoading(false);
      onRefresh();
    });
  };

  const onCheckDappsterOS = (vmid: number) => {
    setLoading(true);
    dispatch(checkDappster(vmid)).then((res) => {
      setLoading(false);
      // onRefresh();
    });
  };
  const confirm: PopconfirmProps["onConfirm"] = (e) => {
    console.log(e);
    message.success("Click on Yes");
  };

  const cancel: PopconfirmProps["onCancel"] = (e) => {
    console.log(e);
    // message.error("Click on No");
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

  const onCreateVM = () => {
    setLoading(true);
    dispatch(createVM({})).then((res) => {
      setLoading(false);
      onRefresh();
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
          <Space
            style={{ marginBottom: 16 }}
            className="w-full flex justify-between"
          >
            <h3>Your Machines</h3>
            <div className="w-full flex-1"></div>
            <Button.Group>
              <Button
                icon={<PlusSquareOutlined />}
                key={"create"}
                onClick={onCreateVM}
              >
                Create New VM
              </Button>
              <Button
                icon={<RiRefreshLine />}
                key={"refresh"}
                onClick={onRefresh}
              >
                Reload
              </Button>
            </Button.Group>
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
