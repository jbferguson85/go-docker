import React from "react";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";
import CsvUpload from "./csv-upload";
describe("<CsvUpload />", () => {
  test("it should mount", () => {
    render(
      <CsvUpload
        onFileUpload={function (file: File): void {
          throw new Error("Function not implemented.");
        }}
      />
    );

    const csvUpload = screen.getByTestId("CsvUpload");

    expect(csvUpload).toBeInTheDocument();
  });
});
