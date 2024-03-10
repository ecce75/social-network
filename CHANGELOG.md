3b14e87 (HEAD -> main, origin/main, origin/HEAD) feat: basic frontend notifications implemented, correct posts feeds for profiles
27aa696 SN-11 InviteGroupMember route fix
0d620fd SN-11 group fixes
1f821f7 fix: fixed fetching and displaying comments below posts Jira: SN-11 SN-22
1fabe5e SN-11 Send group request and group creator request approval set up, decline request need to fix + authCheck for all pages +
97b1092 SN-11 Send group request and group creator request approval set up, decline request need to fix + authCheck for all pages +
d33f8c8 SN-24 added different views on group page based on member, non-member and creator. Implemented group request confirmation view, need to fix
74c2fd5 SN-24 groups fixes,
e8b4100 fix: groups tab now correctly fetches and displays groups, clicking on them correctly routes to group page, group page feed fetches group posts Jira: SN-11
c8d341d fix: groups tab now correctly fetches and displays groups, clicking on them correctly routes to group page Jira: SN-11
d55c03c SN-24 groups fixes
4fa9968 brt
2099c53 SN-24 Create group button api connected
9c86a3a .env upload
eef1c7a fix: removed placeholders for main and profile posts feed, profile icon from dropdown now routes correctly Jira: SN-11 SN-21
a31b549 SN-21 Post creation api connected
0dedfee .gitignore
0d2bf82 fixes
f2463cd Delete .vscode directory
98a26bd Friend request decline and resend fix
a7f53d3 SN-43 Image saving to db based on type. Chat, addFriends and ui fixes
e31c19c (origin/notifications, notifications) merge: merged commit: ca4d244 with notification branch latest; BREAKING CHANGES: go.work.sum might be messy from merge Jira: closes -> SN-68 SN-64 ; updates -> SN-51 SN-12
ca4d244 Merge pull request #16 from AntonUrb/main
3dde392 SN-65 | mod: added documentation to group/groupmember/notification related functions
bb3593b feat: likes&dislikes, group event creation notifies members, delete notification endpoint, posts and comments are now fetched with votes Jira: SN-68 SN-64 Co-Authored by: AntonUrban
bd5d4ef Merge remote-tracking branch 'origin/main'
f1c5cd2 SN-31 SN-27 ws setup done, initial chat working, need to fix styling
06fde31 fix: added proper routes for group invitations and requests Jira: SN-64
a929742 refactor: notification system now resides in notificationHandler which other handlers use to send notifications; notifications are now labeled as types: group, friend, post; removed deprecated code and functions; Jira: SN-64 SN-49 Co-Authored by: AntonUrban BREAKING CHANGES: use of .env file in backend and frontend friendslist.tsx
63ae67b merge: merged group notifications with friend notifications feat: dockerfile Jira: SN-64 SN-12 BREAKING CHANGES: notifications table group_id new field and type check 'group' 'friend' 'post'; groups table deleted bool field; group_members id
71b3b74 Merge pull request #15 from AntonUrb/main
e066c58 SN-64 | mod: modified notification functions to fit CRUD good practices | fix: fixed some parameter inaccuracies
c36358a feat: notifications for sending a friend request, accepting a friend request, commenting on a post, added .env file for global url and port Jira: SN-64 BREAKING CHANGES: notifications table new field sender_id
2710c7e mod: refactored function for getting pending group invitations via cookie authorization
272b481 SN-64 | feat: implemented notifications functionality to group related actions
3feef7d SN-63 addFriends modal fixes
f375e2d SN-63 addFriends modal fixes
a4e4765 SN-63 started implementing addFriends modal with friend requests handling + fixes
dc6447d SN-50 SN-21 perf: added group member checks for posts and events BREAKING CHANGE: database.db reset
73d4f85 Merge branch 'main' of https://github.com/ecce75/social-network
68dc53a md
ff311ab Merge pull request #10 from AntonUrb/main
7460664 Merge branch 'main' into main
c4c7ce3 Documentation for Events & Notifications
bdeafae SN-20 feat: endpoint for getting the userlist
7aede06 merge conflict bypass
0da77d1 Merge pull request #14 from ChrisRichV/main
6705b88 Merge branch 'ecce75:main' into main
1086bee Rudimentary friendlist, various visual/bug fixes
f5996d4 SN-62 fix: added profilesetting field to UpdateUserProfile
6d29888 Merge branch 'ecce75:main' into main
90104e4 Windows support fixes
9f294d4 Merge branch 'main' of https://github.com/ecce75/social-network SN-62
ce6d3f9 SN-62 update: add user profile editing functionality docs: updated profile editing endpoint docs refers also to: SN-51 SN-19 SN-20 SN-23 BREAKING CHANGES: users table new field profile('public', 'private')
4dd9d99 Merge branch 'ecce75:main' into main
55323d6 filepath fix
426b38a Merge branch 'main' of https://github.com/ecce75/social-network
4dd9eb4 SN-62 feat: added profile functionality, refer to docs for endpoints docs: updated be docs with profile functionality chore: removed some commented code, removed userid from register response message user profile now has endpoints for retrieving user profile and posts
a3ac103 SN-62 feat: added profile functionality, refer to docs for endpoints docs: updated be docs with profile functionality chore: removed some commented code, removed userid from register response message user profile now has endpoints for retrieving user profile and posts Closes: SN-62
8e15939 SN-9
ff0b472 Merge branch 'main' of https://github.com/ChrisRichV/social-network
8176e1a Group inv/req responsive desgin
e807103 Profile page
1592c83 SN-62 getallusersposts for profile
666cdae fix
ea4daea Merge pull request #13 from ecce75/frontend
e39879c (origin/frontend) Merge branch 'main' into frontend
e8721ba Merge branch 'main' of https://github.com/ChrisRichV/social-network into ChrisRichV-main
4a0a0a0 Merge pull request #12 from ChrisRichV/main
743175f SN-61
c6df61f SN-50
5794fad Merge pull request #11 from ChrisRichV/main
39fa0af Groups page/Various fixes/Directory changes
c131ffd SN-58 SN-18 SN-51 | docs: updated backend md, added friends api documentation, removed todo from router
f2165d7 (origin/friends, friends) SN-18 SN-58 implemented endpoints for friend request, accept, decline, block, unblock, getAllFriends; handler level checks if a pending/declined/accepted/blocked friend request exists
d9f5d36 SN-34
ea6166f Group feed stuff
13dcfda Merge pull request #9 from ecce75/main
7644c4f Merge pull request #8 from ecce75/frontend
0a1a615 Merge pull request #7 from ChrisRichV/main
516cf8f Merge branch 'main' of https://github.com/ChrisRichV/social-network
71881ac Pages/Groups/Fixes
9ce4ada Merge branch 'ecce75:main' into main
5340bbd Merge branch 'ecce75:main' into main
c5b2dc1 SN-34
b6a4487 SN-56 SN-47
3b790b2 Merge pull request #6 from ChrisRichV/main
c3ec06f Merge branch 'ecce75:main' into main
9155121 Post creation and post feed
6ab8344 Merge branch 'ecce75:main' into main
dc673f0 Merge branch 'ChrisRichV-main' into frontend merge frontend
b97284b (origin/ChrisRichV-main) frontend merge
2f1ffef Merge pull request #4 from ChrisRichV/main
3e25d4c SN-9
77f2716 SN-56 refactored most of the handlers, and repository for comment, post, session and user; got rid of the global sqlite db var; auth middleware now deprecated since front end already does the checking; added GetSessionToken to util for retrieving the session token from http request;
e62b9df Header/NavBar
6e51dd3 email/username already taken fix
fc969dc Merge pull request #3 from ChrisRichV/main
15a3ee8 Merge remote-tracking branch 'refs/remotes/origin/main'
c2bfc85 Fixes
d03d38c migration fix
fe2bb77 migration fix
ee8560c Merge pull request #2 from ecce75/groups
ff0e541 (origin/groups, groups) MERGE
831e76d SN-47 SN-48 SN-25 merge to main for refactoring; search (TODO:) comments in code for group related todos regarding notifications, routes and future implementations;
b135f86 logout fix
6fc2a1b SN-48 removed deprecated code, finished some TODO marks, reorganized code for better readability; SN-47 removed deprecated code; SN-51 created fresh readme with better structure which now should include new endpoints; SN-25 new endpoints in router.go for new methods
d3d3b3f SN-48 SN-47 SN-25 created middleware for checking authentication for functions that require the user to be logged in; separated concerns so groupHandler/Repository is only for actions with group (remove deprecated code next commit), groupMemberHandler/Repository now deals with group invitations/requests and approvals/denials - lots of commented out deprecated code, will remove in next commit; GroupInvitations values changed: now has JoinUserId and InviteUserId;
aaacc0b SN-48 SN-25 moved group member invite/request & approve/decline functionalities to separate handler and repository to have better readability; added targetUserId to groupInvitations table but currently is not quite implemented in the repository level; groupMemberHandler needs refactor as some of the functionalities are separated into different handlers, need to implement auth checking to the groupMemberHandler.go
02b2c5d Frontend auth working with backend
e7c1ca0 SN-25 SN-48 SN-47  renamed invitations to group_invitations, moved all group related invitations/requests functions on handler and repository layer to  groupHandler.go and groupRepository.go -- needs to be checked if everything works expectedly; fixed group_invitations table
bee1ea1 SN-56 SN-47 SN-25 SN-48 router.go: implemented repository pattern for groups and invitations; db: groups table +updated_at, invitations up and down; grouphandler & invitationHandler and respective repository layer uses repository pattern; server.go: router func now takes db as input;
b3dbedc image upload saves image in local dir
051da90 SN-47 group basic crud, docs update, group members table on delete cascade
5180892 Merge branch 'main' of https://github.com/ecce75/social-network Login and Register integration with backend APIs
041478e SN-11 Login and Register working
c331f54 Merge pull request #1 from h2ving/main
74d29df SN-54 backend docs changes
941ad4d backend documentation upgrade
6e724fd Merge branch 'frontend'
ca692d1 Login and Register page
3aebe8d (origin/backflip, backflip) event, friend, group, invitation, notification handlers skeletons and router.go routes for prementioned functionalities, little fixes in postrepository
887ab8b (origin/temp, temp) comments crud, posts crud, events table, router updates, documentation updates
ca20cab errors
97c6c90 post crud, comment crud partial, fixes to db schema
1ceb6e7 yo
db18488 initial build
76c36ff md fixes
d6675b5 login,logout,sessions,createpost,documentation
49385b3 small fixes to user struct and sql model
a8953d5 structure
b8fd05f migrations working
bed3ba0 be setup
0f7c363 initial build
