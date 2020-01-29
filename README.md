# MakingProgramUsingGo
this repository is to make program what we need using go 

### TODO
- [ ] Install go on your computer
- [ ] Run the Hello World example
- [ ] Upload a screenshot to this repository

### Refer
- [ ] https://www.youtube.com/user/dforensics

### Question
- When there are few duplicate files during the program function, the files can be deleted based on the creation date. However, if there are many duplicate files, the files created late will remain.
  * When checking for duplicate files, use the hash value instead of the creation date.
  
- subong: when he delete files using hash value comparing, it make remain copied files. ex) subong, subong(1), subong(2) ==> subong(2)
        but the reviewer want to see "subong" file. because it might creat first possibly.
        So what he want to fix his program is that compare hash and delete duplicate file. but remain first files.
        Could you please advise these things?
