import React, { useState } from 'react'
import { connect } from 'react-redux'
import { addEntry } from './journalSlice'

const mapDispatch = { addEntry }

const Journal = ({ addEntry }) => {
  const [entryTitle, setEntryTitle] = useState('')
  const [entryBody, setEntryBody] = useState('')

  const onTitleChange = e => setEntryTitle(e.target.value)
  const onBodyChange = e => setEntryBody(e.target.value)

  return (
    <div>
      <form
        onSubmit={e => {
          e.preventDefault()
          if (!entryTitle.trim()) {
            return
          }
          addEntry(entryTitle, entryBody.trim())
          setEntryTitle('')
          setEntryBody('')
        }}
      >
        <input value={entryTitle} onChange={onTitleChange} />
        <input value={entryBody} onChange={onBodyChange} />
        <button type="submit">Add Entry</button>
      </form>
    </div>
  )

}

export default connect(null, mapDispatch)(Journal)
