export interface NodeData {
  Status: string;
  ID: string;
  Node: string;
  Type: string;
  MaxCPU: number;
  MaxMem: number;
  Disk: number;
  ssl_fingerprint: string;
  MaxDisk: number;
  Mem: number;
  CPU: number;
  Uptime: number;
}
export interface VMData {
  VirtualMachineConfig: any;
  Name: string;
  Node: string;
  Agent: number;
  NetIn: number;
  CPUs: number;
  DiskWrite: number;
  Status: string;
  VMID: number;
  PID: number;
  Netout: number;
  Disk: number;
  Mem: number;
  CPU: number;
  MaxMem: number;
  MaxDisk: number;
  DiskRead: number;
  Uptime: number;
  Template: boolean;
  HA: {
    Managed: number;
  };
  tags: string;
}
