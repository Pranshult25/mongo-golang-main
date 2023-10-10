import {useEffect, useState} from "react";
import axios from "axios";
import Post from "./Post";
import { Alert } from "bootstrap";

function SearchResultsPage(props) {
  const {text} = props.match.params;
  const [comments,setComments] = useState([]);
  let postClasses = "block border rounded-md bg-reddit_dark-brighter p-3 mx-6 border-2 border-reddit_border text-reddit_text pb-4";
  useEffect(() => {
    axios.get('http://localhost:4000/comments', { withCredentials: true })
      .then(response => {
        const commentdata = response.data;
        const filteredComments = commentdata.filter(comment => (
          comment.body.includes(text) || comment.title.includes(text)
        ));
        const reversedComments = [...filteredComments].reverse();
        setComments(reversedComments);
      })
      .catch(error => {
        console.error('Error fetching comments:', error);
      });
  }, [text]);


  return (
    <div className="bg-reddit_dark">
      {comments.length > 0 ? (
        comments.map(comment => (
          <Post {...comment} isListing={true} />
        ))
      ) : (
        <div>
          {/* Display an alert or message here */}
          <div className={postClasses}>
            <h2 className="text-xl mb-3">No content found with "{text}"</h2>
        </div>
          {/* Redirect to "/" */}
          
        </div>
      )}
    </div>
  );
  
}

export default SearchResultsPage;