import { Link } from 'react-router-dom'
import React from 'react'

export default function Home() {
    return (
        <div>

            <button id="dropdownSearchButton" data-dropdown-toggle="dropdownSearch" data-dropdown-placement="bottom" style={{border:"1px solid #313638"}} class="text-white bg-black-700 hover:bg-black-800 font-medium rounded-lg text-sm px-1.5 py-1 text-center inline-flex items-center" type="button"><svg style={{ marginRight: "22px", marginLeft: "0px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">

                <path strokeLinecap="round" strokeLinejoin="round" d="M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" />
            </svg>
                Home
                <svg style={{ marginLeft: "150px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                    <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                </svg>
            </button>


            <div id="dropdownSearch" style={{backgroundColor:"#313638"}} class="z-10 hidden rounded-lg shadow w-60">
                <div class="p-3">
                    <label for="input-group-search" class="sr-only">Search</label>
                    <div class="relative">

                        <input style={{ backgroundColor: "#272729", borderColor: "1px solid #343536", marginTop: "13px" }} type="text" id="input-group-search" class="block w-full p-2 pl-3 text-sm text-white border rounded-lg bg-gray-50 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Filter" />
                    </div>
                </div>


                <div style={{ color: "rgb(129, 131, 132)", paddingLeft: "15px", fontSize: "15px", marginTop: "10px"}}>Feeds</div>
                <ul class="py-2 text-sm text-white dark:text-gray-200" aria-labelledby="multiLevelDropdownButton">
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" />
                        </svg>
                        <Link to="/" class="block px-2 py-2 hover:bg-black">Home</Link>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M2.25 18L9 11.25l4.306 4.307a11.95 11.95 0 015.814-5.519l2.74-1.22m0 0l-5.94-2.28m5.94 2.28l-2.28 5.941" />
                        </svg>
                        <a href="#" class="block px-4 py-2 hover:bg-black">Popular</a>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z" />
                        </svg>

                        <a href="#" class="block px-4 py-2 hover:bg-black">All</a>
                    </li>



                    <div style={{ color: "rgb(129, 131, 132)", paddingLeft: "15px", fontSize: "15px", marginTop: "10px" }}>Other</div>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "13px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 6h9.75M10.5 6a1.5 1.5 0 11-3 0m3 0a1.5 1.5 0 10-3 0M3.75 6H7.5m3 12h9.75m-9.75 0a1.5 1.5 0 01-3 0m3 0a1.5 1.5 0 00-3 0m-3.75 0H7.5m9-6h3.75m-3.75 0a1.5 1.5 0 01-3 0m3 0a1.5 1.5 0 00-3 0m-9.75 0h9.75" />
                        </svg>

                        <button id="doubleDropdownButton" data-dropdown-toggle="doubleDropdown" data-dropdown-placement="right-start" type="button" class="flex items-center justify-between w-full px-4 py-2 hover:bg-black dark:hover:bg-black-600 dark:hover:text-white">Categories<svg class="w-2.5 h-2.5 ml-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4" />
                        </svg></button>
                        <div id="doubleDropdown" style={{ backgroundColor: "#313638" }} class="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700">
                            <ul class="py-2 text-sm text-white dark:text-gray-200" aria-labelledby="doubleDropdownButton">
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/Academics" class="block px-4 py-2 hover:bg-black">Academics</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/Technology" class="block px-4 py-2 hover:bg-black">Technology</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/Health" class="block px-4 py-2 hover:bg-black">Health</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/TravelAndLeisure" class="block px-4 py-2 hover:bg-black">Travel and Leisure</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/CareerAndEmployment" class="block px-4 py-2 hover:bg-black">Career and Employment</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/LifestyleAndRelationship" class="block px-4 py-2 hover:bg-black">Lifestyle and Relationship</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/FinanceAndMoney" class="block px-4 py-2 hover:bg-black">Finance and Money</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/RelationshipAndSocial" class="block px-4 py-2 hover:bg-black">Relationship and Social</a>
                                </li>
                                <li className='hover:bg-black'>
                                    <a href="/commentsbycategory/HobbiesAndIntrests" class="block px-4 py-2 hover:bg-black">Hobbies and Intrests</a>
                                </li>
                               
                            </ul>
                        </div>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z" />
                        </svg>
                        <a href="#" class="block px-4 py-2 hover:bg-black">User Settings</a>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M7.5 8.25h9m-9 3H12m-9.75 1.51c0 1.6 1.123 2.994 2.707 3.227 1.129.166 2.27.293 3.423.379.35.026.67.21.865.501L12 21l2.755-4.133a1.14 1.14 0 01.865-.501 48.172 48.172 0 003.423-.379c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0012 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741v6.018z" />
                        </svg>
                        <a href="#" class="block px-4 py-2 hover:bg-black">Messages</a>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
                        </svg>
                        <a href="#" class="block px-4 py-2 hover:bg-black">Create Post</a>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0" />
                        </svg>
                        <a href="#" class="block px-4 py-2 hover:bg-black">Notifications</a>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M12 6v12m-3-2.818l.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <a href="#" class="block px-4 py-2 hover:bg-black">Premium</a>
                    </li>
                    <li style={{ display: "flex", alignItems: "center" }} className='hover:bg-black'>
                        <svg style={{ marginLeft: "10px" }} xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" />
                        </svg>
                        <a href="#" class="block px-4 py-2 hover:bg-black">Avatar</a>
                    </li>
                </ul>
            </div>

        </div>
    )
}
