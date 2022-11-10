import { NextComponentType } from "next";


const Input = ({placeholder}:{placeholder:string}) => {
    return (
        <>
            <input type="text" placeholder={placeholder} className={`
                border
                border-sky-300 
                w-full 
                p-2 
                rounded-lg 
                px-3
                bg-slate-900
            `} />
        </>
    )
}

export default Input