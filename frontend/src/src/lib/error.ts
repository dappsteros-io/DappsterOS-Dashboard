export class JsonError extends Error {
  constructor(message: string, data: any) {
    super(message);
    this.name = "JsonError";
    this.data = data;
  }
  data?: any;
}
