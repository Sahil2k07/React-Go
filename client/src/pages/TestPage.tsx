import reactLogo from "../assets/react.svg";
import viteLogo from "../../public/vite.svg";
import { useState } from "react";
import exampleService from "../services/exampleService";

function TestPage() {
  const [response, setResponse] = useState<string>("");

  async function handleClick() {
    const { data } = await exampleService.testExampleAPI();
    console.log(data);
    setResponse(data);
  }
  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank" rel="noreferrer">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank" rel="noreferrer">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={handleClick}>Make API Call</button>
        <p>API Response: {response}</p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default TestPage;
