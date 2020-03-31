import React from "react";

const Entry = ({ entry }) => (
  <li
    className="entry-item"
    onClick={() => {} /** dispatches action to toggle todo */}
  >
    <span>
      {entry.title}
    </span>
  </li>
);

export default Entry;