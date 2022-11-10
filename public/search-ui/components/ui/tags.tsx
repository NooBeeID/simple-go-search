
const Tags = ({name}:{name:string}) => {
    return (
       <div className="cursor-pointer hover:bg-sky-600 hover:text-bold hover:text-slate-900 border border-sky-300 p-2 text-center text-sky-300 rounded-lg">
            <div className="flex justify-between gap-3">
                <span>{name}</span>
                <span className="rounded-md bg-sky-300 text-slate-900 px-1">10</span>
            </div>
       </div>
    )
}

export default Tags