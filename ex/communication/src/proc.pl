#!/usr/bin/perl -l
my (%child, $pid, $parent);
my @ps=`ps -e -opid,ppid,comm`;	  # capture the output from `ps`
foreach (@ps[1..$#ps]) {	  # discard the header line
  ($pid, $parent, undef) = split; # split the line, discard 'comm'
  push @{$child{$parent}}, $pid;  # save the child PIDs on a list
}
# Walk through the sorted PPIDs
foreach (sort { $a <=> $b } keys %child) {  
  print "Pid ", $_, " has ", @{$child{$_}}+0, " child",
    @{$child{$_}} == 1 ? ": " : "ren: ", "[@{$child{$_}}]";
}
