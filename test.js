lang = [
    {id: 1, lang: "en"}, {id: 2, lang: "id"}
]

inp = [
    {name: "inp-1", "inp": "Array Indo"}, {name: "inp-2", "inp": "Array English"},
]

newArr = []
let check = lang.map(e => {
    inp.map(ele => {
        ele.name.includes(e.id.toString()) ? newArr.push({language: e.id, inp: ele.inp}) : ""
    })
})

console.log(newArr)