import { NextPage } from "next"
import Input from "../components/form/input"
import Card from "../components/ui/card"
import Tags from "../components/ui/tags"


const Home: NextPage = () => {
  return (
   <div className="container max-w-screen-lg m-auto mt-10 p-3">
      <div className="grid grid-cols-6 gap-4">
        <div className="filter col-span-2 col-start-1">
          <div className="border border-sky-300 p-3 bg-slate-900 rounded-lg">
            <span className="text-sm">Filter</span>
            <hr />
            <div className="grid grid-cols-2 gap-4 mt-5 text-sm ">
              <h1 className="col-span-2 font-semibold">Category</h1>
              <Tags name="Frontend"/>
              <Tags name="Backend"/>
              <Tags name="UI / UX"/>
              <Tags name="DevOps"/>
            </div>
          </div>
        </div>
        <div className="search  col-span-4 col-start-3 ">
          <Input placeholder="Search ..."/>

          <div className="content">
            <div className="grid grid-cols-2 gap-3">
              <Card 
                imgSrc="https://academy.alterra.id/blog/wp-content/uploads/2021/07/golang-img.png"
                title="Belajar Golang"
              />
              <Card 
                title="Belajar Javascript"
                imgSrc="https://res.cloudinary.com/noobeeid/image/upload/q_60/v1638689477/Courses/ReactJS_d0yqoq.jpg"
                />
            </div>
          </div>
        </div>
      </div>
   </div>
  )
}

export default Home