import Header from "./Header";
import { BrowserRouter as Router, Redirect, Route, Switch } from "react-router-dom";
import RoutingSwitch from "./RoutingSwitch";
import PostFormModal from "./PostFormModal";
import AuthModal from "./AuthModal";
import { useContext, useEffect , useState } from "react";
import RedirectContext from "./RedirectContext";
import UserContext from "./UserContext"; // Import the UserContext
import ProfilePage from "./ProfilePage";



function Routing() {
  const { redirect, setRedirect } = useContext(RedirectContext);
  const user = useContext(UserContext); // Get user context
  const [showProfileModal, setShowProfileModal] = useState(false); // State to control profile modal

  useEffect(() => {
    if (redirect) {
      setRedirect(false);
    }
  }, [redirect]);

return (
  <Router>
    {redirect && <Redirect to={redirect} />}
    {!redirect && (
      <>
        <Header setShowProfileModal={setShowProfileModal} />
        <RoutingSwitch />
        <PostFormModal />
        <AuthModal />
        
        {/* Show profile page if user is logged in */}
        {user.username && (
          <Route path="/profile">
            <ProfilePage />
          </Route>
        )}
      </>
    )}
  </Router>
);
}

export default Routing;