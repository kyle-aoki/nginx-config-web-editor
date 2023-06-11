// ############################################################################
const LogInPath = "/log-in";
interface LogInInput {
  Username: string;
  Password: string;
}
interface LogInOutput {
  Successful: boolean;
  Session: string;
}

// ############################################################################
const UserFindPath = "/user/find";
interface UserFindOutput {
  Username: string;
}

// ############################################################################
const NginxReloadPath = "/nginx/reload";
interface NginxReloadInput {
  Name: string;
}
interface NginxReloadOutput {
  Result: string;
  Error: string;
}

// ############################################################################
const NginxClonePath = "/nginx/clone";
interface NginxCloneInput {
  Name: string;
}
interface NginxCloneOutput {
  NewName: string;
}

// ############################################################################
const NginxRenamePath = "/nginx/rename";
interface NginxRenameInput {
  Name: string;
  NewName: string;
}

// ############################################################################
const NginxListPath = "/nginx/list";
interface NginxListOutput {
  Files: string[];
}

// ############################################################################
const NginxReadPath = "/nginx/read";
interface NginxReadInput {
  Name: string;
}
interface NginxReadOutput {
  Value: string;
}

// ############################################################################
const NginxSavePath = "/nginx/save";
interface NginxSaveInput {
  Name: string;
  Value: string;
}

// ############################################################################
const NginxDeletePath = "/nginx/delete";
interface NginxDeleteInput {
  Name: string;
}

// ############################################################################
const NginxTestPath = "/nginx/test";
interface NginxTestInput {
  Name: string;
}
interface NginxTestOutput {
  Result: string;
  Error: string;
}

// ############################################################################

export class Call<T> {
  output: T;
  error: string | null;
  constructor(output: T, error: string | null) {
    this.output = output;
    this.error = error;
  }
}

// prettier-ignore
export default class NcweClient {
  host: string;
  private session: string = localStorage.getItem("session") || "";
  constructor(host: string) {
    this.host = host;
  }
  private async rq<T>(path: string, body: any): Promise<Call<T>> {
    const resp = await fetch(`${this.host}${path}`, {
      method: "POST",
      headers: { session: this.session },
      body: JSON.stringify(body)
    });
    if (resp.status !== 200) {
      return new Call(null as T, resp.statusText)
    }
    const contentlength = resp.headers.get("content-length");
    if (contentlength !== null && contentlength === "0") return new Call(null as T, null);
    return new Call(await resp.json(), null);
  }
  setSession  = (session: string)                                                 => { this.session = session; };
  LogIn       = async (input: LogInInput): Promise<Call<LogInOutput>>             => this.rq(LogInPath, input);
  UserFind    = async (): Promise<Call<UserFindOutput>>                           => this.rq(UserFindPath, null);
  NginxReload = async (input: NginxReloadInput): Promise<Call<NginxReloadOutput>> => this.rq(NginxReloadPath, input);
  NginxClone  = async (input: NginxCloneInput): Promise<Call<NginxCloneOutput>>   => this.rq(NginxClonePath, input);
  NginxRename = async (input: NginxRenameInput)                                   => this.rq(NginxRenamePath, input);
  NginxList   = async (): Promise<Call<NginxListOutput>>                          => this.rq(NginxListPath, null);  
  NginxRead   = async (input: NginxReadInput): Promise<Call<NginxReadOutput>>     => this.rq(NginxReadPath, input);
  NginxSave   = async (input: NginxSaveInput)                                     => this.rq(NginxSavePath, input);
  NginxDelete = async (input: NginxDeleteInput)                                   => this.rq(NginxDeletePath, input);
  NginxTest   = async (input: NginxTestInput): Promise<Call<NginxTestOutput>>     => this.rq(NginxTestPath, input);
}

export const ncweClient = new NcweClient("http://18.191.251.159:9040");
