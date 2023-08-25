alter table book_managment
   add constraint inlab_nonnegative check (total_in_lab >= 0);

alter table book_managment
   add constraint outlab_nonnegative check (total_out_lab >= 0);

