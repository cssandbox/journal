import React from "react";

export default function Entries(props) {
  const entries = props.value;
  const listItems = entries.map((entry) =>
    <li key={entry.uuid}>{entry.title}</li>
  );
  return (
    <div>
      <div>Add New</div>
      <ul>{listItems}</ul>
    </div>
  );
}