import { useState, useEffect } from "react";
import axios from "axios";

function App() {
  const [message, setMessage] = useState("");
  const instance = axios.create({ baseURL: "http://localhost:8080" });

  useEffect(() => {
    instance
      .get("/api/message")
      .then((response) => setMessage(response.data.text))
      .catch((error) => console.error(error));
  }, []);

  return (
    <div>
      <h1>{message}</h1>
    </div>
  );
}

export default App;
