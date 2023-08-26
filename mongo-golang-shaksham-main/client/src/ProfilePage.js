import React, { useContext, useState, useEffect } from "react";
import UserContext from "./UserContext";
import axios from "axios";
import Comment from "./Comment"; 

function ProfilePage() {
  const user = useContext(UserContext);
  const [userComments, setUserComments] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (user.username) {
      // Fetch user's past comments
      axios.get(`http://localhost:4000/comments/user/${user.username}`, { withCredentials: true })
        .then(response => {
          setUserComments(response.data);
          setIsLoading(false);
        })
        .catch(error => {
          console.error("Error fetching user's comments:", error);
          setIsLoading(false);
        });
    } else {
      setIsLoading(false);
    }
  }, [user.username]);

  if (isLoading) {
    return (
      <div className="h-screen bg-reddit_dark text-reddit_text flex items-center justify-center">
        <p>Loading...</p>
      </div>
    );
  }

  return (
    <div className="h-screen bg-reddit_dark text-reddit_text overflow-hidden">
      <div className="max-w-screen-xl mx-auto py-10 px-4 h-full flex flex-col justify-between">
        {user.username ? (
          <div>
            <div className="flex items-center mb-5">
               <img src={user.avatar} alt="Profile Avatar" className="w-12 h-12 rounded-full mr-3" />
              <h1 className="text-2xl font-semibold">Name: {user.username}</h1>
            </div>
            
            <div className="flex-grow mb-8 overflow-y-auto scrollbar-thin scrollbar-thumb-reddit_dark-lighter scrollbar-track-reddit_dark">
              <h2 className="text-lg mb-3">Your Past Queries</h2>
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
          </div>
        ) : (
          <div className="h-full flex items-center justify-center">
            <div className="text-center">
              <img src="path/to/error-image.png" alt="Error" className="w-20 h-20 mx-auto mb-4" />
              <p className="text-lg">Oops! Please log in to access your profile.</p>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default ProfilePage;
