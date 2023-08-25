import React, { useContext, useState, useEffect } from "react";
import UserContext from "./UserContext";
import axios from "axios";
import Comment from "./Comment"; 

function ProfileModal({ onClose }) {
  const user = useContext(UserContext);
  const [userComments, setUserComments] = useState([]);

  useEffect(() => {
    if(!user.username) return; // If user is not logged in, do nothing (don't fetch comments
    // Fetch user's past comments
    axios.get(`http://localhost:4000/comments/user/${user.username}`, { withCredentials: true })
      .then(response => {
        setUserComments(response.data);
      })
      .catch(error => {
        console.error("Error fetching user's comments:", error);
      });
  }, [user.username]);

  return (
    <div className="w-screen h-screen fixed top-0 left-0 z-20 flex" style={{ backgroundColor: "rgba(0,0,0,.8)" }}>
      <div className="border border-reddit_dark-brightest w-4/5 md:w-3/5 lg:w-2/3 xl:w-1/2 bg-reddit_dark p-5 text-reddit_text self-center mx-auto rounded-md">
        <h1 className="text-2xl mb-5">Profile: {user.username}</h1>

        <h2 className="text-lg mb-3">Your Past Comments</h2>
        <div className="max-h-72 overflow-y-auto scrollbar-thin scrollbar-thumb-reddit_dark-lighter scrollbar-track-reddit_dark">
          {userComments.length > 0 ? (
            <div>
              {userComments.map(comment => (
                <Comment key={comment._id} comment={comment} />
              ))}
            </div>
          ) : (
            <p>No comments yet.</p>
          )}
        </div>

        <button onClick={onClose} className="px-4 py-2 mt-4 bg-reddit_dark-brightest text-reddit_text rounded-md">
          Close
        </button>
      </div>
    </div>
  );
}

export default ProfileModal;
