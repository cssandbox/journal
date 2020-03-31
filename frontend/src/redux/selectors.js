export const getJournalState = store => store.journal

export const getJournalList = store =>
    getJournalState(store) ? getJournalState(store).allIds : []

export const getEntryById = (store, id) =>
getJournalState(store) ? { ...getJournalState(store).byIds[id], id } : {}

export const getEntries = store =>
    getJournalList(store).map(id => getEntryById(store, id))
