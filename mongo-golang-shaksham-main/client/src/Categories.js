import React, {useState} from 'react'

const Categories = (props) => {
    const [name, setName] = useState("Categories")
    

    const changeName = (e) =>{
        setName(e.target.innerText)
        props.setCategory(e.target.innerText)
    }
    return (

        <>
            <button id="dropdownDefaultButton" data-dropdown-toggle="dropdown" class="hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-2 py-2.5 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700" style={{ backgroundColor: "rgb(26, 26, 26)", color: "rgb(215, 218, 220)"}} type="button">{name} <svg class="w-2.5 h-2.5 ml-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 10 6">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 4 4 4-4" />
            </svg></button>
            <div id="dropdown" class="z-10 hidden">
            <ul class="py-2 text-sm text-gray-700 dark:text-gray-200 flex-grow" aria-labelledby="dropdownDefaultButton" style={{backgroundColor: "rgb(26, 26, 26)", color:"white"}}>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-black-600 dark:hover:text-black" onClick={changeName}>Academics</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-black-600 dark:hover:text-black" onClick={changeName}>Technology</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-white-600 dark:hover:text-black" onClick={changeName}>Health</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-white-600 dark:hover:text-black" onClick={changeName}>Travel and Leisure</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-white-600 dark:hover:text-black" onClick={changeName}>Career and Employment</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-white-600 dark:hover:text-black" onClick={changeName}>Lifestyle and Relationships</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-white-600 dark:hover:text-black" onClick={changeName}>Finance and Money</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-white-600 dark:hover:text-black" onClick={changeName}>Relationships and Social</a>
                    </li>
                    <li>
                        <a class="block px-4 py-2 hover:bg-gray-100 hover:text-black dark:hover:bg-white-600 dark:hover:text-black" onClick={changeName}>Hobbies and Interests</a>
                    </li>
                </ul> 
            </div>


        </>
    )
}

export default Categories
