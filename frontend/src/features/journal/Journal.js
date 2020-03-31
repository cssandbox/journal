import React, { useState } from 'react'
import { connect } from 'react-redux'
import { addEntry } from './journalSlice'

const mapDispatch = { addEntry }

const Journal = ({ addEntry }) => {
  const [entryTitle, setEntryTitle] = useState('')

  const onChange = e => setEntryTitle(e.target.value)

  return (
    <div>
      <form
        onSubmit={e => {
          e.preventDefault()
          if (!entryTitle.trim()) {
            return
          }
          addEntry(entryTitle)
          setEntryTitle('')
        }}
      >
        <input value={entryTitle} onChange={onChange} />
        <button type="submit">Add Entry</button>
      </form>
    </div>
  )

}

export default connect(null, mapDispatch)(Journal)
