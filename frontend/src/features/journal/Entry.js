import React from "react";

const Entry = ({ entry }) => (
  <li
    className="entry-item"
    onClick={() => {} /** dispatches action to toggle todo */}
  >
    <span>
      {entry.title} | {entry.body}
    </span>
  </li>
);

export default Entry;