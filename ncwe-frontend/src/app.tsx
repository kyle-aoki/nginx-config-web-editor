import { Editor } from "@monaco-editor/react";
import { useEffect, useState } from "react";
import { ncweClient } from "./nwce-client";
import { InfoPopupText } from "./popup";
import {
  AppContainer,
  ControlBarContainer,
  ControlBarLeftContainer,
  Inner,
  Outer,
  TextEditorContainer,
  Text,
  Button,
  ControlBarRightContainer,
} from "./styled-components";
import { useAsyncEffect } from "./util";

export const App = () => {
  const [files, setFiles] = useState<string[]>([]);
  const [file, setFile] = useState<string>("");
  const [content, setContent] = useState<string>("");
  const [edited, setEdited] = useState<boolean>(false);
  const [error, setError] = useState<string>("");
  const [status, setStatus] = useState<string>("");

  useEffect(() => {
    if (error.length !== 0) alert(error);
  }, [error]);

  useEffect(() => {
    if (status.length === 0) return;
    setTimeout(() => setStatus(""), 5000);
  }, [status]);

  useAsyncEffect(async () => {
    if (files.length === 0) return;
    const nro = await ncweClient.NginxRead({ Name: file });
    setContent(nro.output.Value);
  }, [file]);

  useAsyncEffect(async () => {
    await fillSelect();
  }, []);

  const fillSelect = async () => {
    const files = (await ncweClient.NginxList()).output.Files;
    setFiles(files);
    if (files.length > 0) {
      setFile(files[0]);
    }
  };

  const handleSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setFile(e.target.value);
  };

  const handleTextEditorChange = (str: string | undefined) => {
    setEdited(true);
    setContent(str || "");
  };

  const handleDiscard = async () => {
    const nro = await ncweClient.NginxRead({ Name: file });
    setContent(nro.output.Value);
    setEdited(false);
  };

  const handleSave = async () => {
    await ncweClient.NginxSave({ Name: file, Value: content });
    setEdited(false);
  };

  const handleRename = async () => {
    const newName = prompt("enter new name:");
    if (newName === null) return;
    await ncweClient.NginxRename({ Name: file, NewName: newName });
    await fillSelect();
    setFile(newName);
  };

  const handleClone = async () => {
    const nco = await ncweClient.NginxClone({ Name: file });
    const nlo = await ncweClient.NginxList();
    setFiles(nlo.output.Files);
    setFile(nco.output.NewName);
  };

  const handleDelete = async () => {
    if (confirm(`delete ${file}?`)) {
      await ncweClient.NginxDelete({ Name: file });
      setFiles((await ncweClient.NginxList()).output.Files);
      setFile(files[0]);
    }
  };

  const handleReload = async () => {
    const nro = await ncweClient.NginxReload({ Name: file });
    if (nro.output.Error) {
      setError(nro.output.Error);
      return;
    }
    setStatus("reload successful");
  };

  const handleTest = async () => {
    const nto = await ncweClient.NginxTest({ Name: file });
    if (nto.output.Error) {
      setError(nto.output.Error);
      return;
    }
    setStatus("conf file is valid");
  };

  return (
    <>
      <AppContainer>
        <ControlBarContainer>
          <ControlBarLeftContainer>
            <select onChange={handleSelect} value={file} style={{ minWidth: "200px" }}>
              {files.map((file, idx) => {
                return <option key={idx}>{file}</option>;
              })}
            </select>
            <Button onClick={handleClone}>clone</Button>
            <Button onClick={handleDelete}>delete</Button>
            <Button onClick={handleRename}>rename</Button>
            {edited && (
              <>
                <Text style={{ width: "15px" }}>|</Text>
                <Button style={{ width: "75px" }} onClick={handleDiscard}>
                  discard
                </Button>
                <Button style={{ width: "75px" }} onClick={handleSave}>
                  save
                </Button>
              </>
            )}
          </ControlBarLeftContainer>
          <ControlBarRightContainer>
            <Button
              onClick={() => {
                localStorage.clear();
                location.reload();
              }}
            >
              log-out
            </Button>
            <Button onClick={() => alert(InfoPopupText)}>info</Button>
            <Button onClick={handleReload}>reload</Button>
            <Button onClick={handleTest}>test</Button>
            <Text>{status}</Text>
          </ControlBarRightContainer>
        </ControlBarContainer>
        <TextEditorContainer>
          <Outer>
            <Inner>
              <Editor
                onChange={handleTextEditorChange}
                theme="vs-dark"
                language="html"
                value={content}
                options={{ readOnly: false }}
              />
            </Inner>
          </Outer>
        </TextEditorContainer>
      </AppContainer>
    </>
  );
};
