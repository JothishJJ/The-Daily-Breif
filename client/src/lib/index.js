/**
 * @param {any} description 
 *
 * @returns string 
*/
export function getSliceDesc(description) {
    const length = description.indexOf(">")
    const sliceText = description.slice(length + 1)

    return sliceText
}
