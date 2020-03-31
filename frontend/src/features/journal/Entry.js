import React from "react";


export default function Entry(props) {
  const entry = props.value;
  return (
    <div>
      <div> Back </div>
      <div>{entry.title}</div>
    </div>
  );
}
