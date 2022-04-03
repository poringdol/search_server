function capitalize(str) {
    if (str === undefined) {
        return str
    }
    return str.replace(/(^|\s)\S/g, function (a) {
        return a.toUpperCase()
    })
}

export default capitalize;