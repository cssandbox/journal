import React from "react";
import { connect } from 'react-redux'
import { addEntry } from '../redux/actions'

class Journal extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      input: ''
    }
  }

  updateInput = input => {
    this.setState({ input });
  };

  handleAddEntry = () => {
    // dispatches actions to add todo
    this.props.addEntry(this.state.input)

    // sets state back to empty string
    this.setState({ input: '' })
  }

  render() {
    return (
      <div>
        <input
          onChange={e => this.updateInput(e.target.value)}
          value={this.state.input}
        />
        <button className="add-entry" onClick={this.handleAddEntry}>
          Add Entry
        </button>
      </div>
    )
  }
}

export default connect(
  null,
  { addEntry }
)(Journal)
