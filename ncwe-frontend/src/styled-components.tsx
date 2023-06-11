import styled, { createGlobalStyle } from "styled-components";

export const GlobalStyles = createGlobalStyle`
  html, body, #root {
    margin: 0;
    background-color: #171717;
  }
  * {
    box-sizing: border-box;
  }
`;

export const AppContainer = styled.div`
  width: 100vw;
  height: 100vh;
`;

export const ControlBarLeftContainer = styled.div`
  display: flex;
  gap: 10px;
`;

export const ControlBarRightContainer = styled.div`
  display: flex;
  flex-direction: row-reverse;
  gap: 10px;
  margin-left: auto;
`;

export const ControlBarContainer = styled.div`
  height: 50px;
  width: 100%;
  background-color: #121212;
  border-bottom: 1px solid #1f1f1f;
  display: flex;
  align-items: center;
  padding: 0px 20px;
  gap: 10px;
`;

export const TextEditorContainer = styled.div`
  height: calc(100% - 50px);
`;

export const Outer = styled.div`
  height: 100%;
  width: 100%;
  max-height: 100%;
  max-width: 100%;
  overflow: hidden;
  position: relative;
`;
export const Inner = styled.div`
  height: 100%;
  width: 100%;
  left: 0;
  top: 0;
  overflow: hidden;
  position: absolute;
`;

export const Text = styled.span`
  white-space: nowrap;
  overflow: hidden;
  color: white;
  font-family: monospace;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
`;

export const Button = styled.button`
  min-width: 60px;
`;
