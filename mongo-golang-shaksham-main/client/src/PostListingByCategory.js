import {useState,useEffect} from "react";
import axios from "axios";
import Post from "./Post";

function PostsListingByCategory(props) {

  const [comments,setComments] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:4000/commentsbycategory/' + props.category, {withCredentials:true})
      .then(response => setComments(response.data));

  }, []);


  return (
    <div className="bg-reddit_dark">
      {comments.map(comment => (
        <Post {...comment} isListing={true} />
      ))}
    </div>
  );
}

export default PostsListingByCategory;