import Tags from "./tags";


const Card = ({imgSrc, title}:{imgSrc:string, title:string}) => {
    return (
        <div className="card mt-3 border border-slate-500 bg-slate-900 rounded-lg">
            <div className="img-banner h-64 overflow-y-hidden rounded-lg">
                {
                    imgSrc !== ""
                        ? <div className="img"
                            style={{
                                backgroundImage:`url(${imgSrc})`,
                                backgroundSize:"cover",
                                backgroundPosition:"center center",
                                height:"100%",
                                width:"100%"
                            }}
                            >
                        </div>
                        : <h1>No Image</h1>
                }
            </div>
            <div className="detail p-3">
                <h1 className=" font-semibold">{title}</h1>
                <span className="text-sm text-slate-400">#Frontend</span>
                <p className="font-bold mb-5">Rp. 500,000</p>
                <button className="bg-sky-300 text-slate-900 w-full p-2 mt-3 rounded">
                    Add to card
                </button>
            </div>
        </div>
    )
}

export default Card;