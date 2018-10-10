#!/usr/bin/python
# -*- coding: utf-8 -*-

""" Implementation of the Square Reciprocal Iterative Filtering voting algorithm """

__author__      = "Charlie Bradford"

import numpy as np
from random import randint
from statistics import mean

VOTERS = 15    # Total number of voters
DISSIDENTS = 2 # Number of dissedents colluding 
ELECTIONS = 10 # Total number of elections
BAD_ELEC  = 5  # The election that dissenters want to rig
ITEMS = 5      # Numbers of items per votes
QUIET = False





def get_vote(votes):
    est = [mean(election) for election in votes]

    rounds = 0

    [print(election) for election in votes]



    diff = 1
    trustworthiness = [1.0 for voter in range(VOTERS)]
    while diff > 0.001:
        distance = [0.0 for voter in range(VOTERS)]
        for i, election in enumerate(votes):
            for j, vote in enumerate(election):
                distance[j] += (vote - est[i])**2


        normalisation = sum([(1.0/dist) for dist in distance])
        print("Normalisation = {}".format(normalisation))
        trustworthiness = [(1.0/distance[i])/normalisation for i in range(VOTERS)]
        for i in range(VOTERS):
            print("Distance for voter {} is {}, trustworthiness is {}".format(i+1, distance[i], trustworthiness[i]))

        old = est
        est = [sum([trustworthiness[i] * vote for vote in election]) for election in votes]
        if not QUIET:
            [print("New estimate for election {} is {}".format(i+1,j)) for i, j in enumerate(est)]

        diff = sum([abs(i - j) for i, j in zip(old, est)])
        rounds += 1
    print("Converged after {} iterations".format(rounds))
    [print("Final estimate for election {} is {}".format(i+1,int(j))) for i, j in enumerate(est)]



    return est

def create_votes():
    votes = []
    truth = []
    for i in range(ELECTIONS):
        elec_votes = []
        true_value = randint(0, ITEMS-1)
        truth.append(true_value)
        fake_value = int(true_value + ITEMS/2) % ITEMS
        if i == BAD_ELEC:
            print("Fake value for election {} is {}".format(i+1, fake_value))
        for voter in range(VOTERS-DISSIDENTS ):
            voter_bias = randint(-1, 1)
            voter_vote = abs(true_value + voter_bias)
            if voter_vote >= ITEMS:
                voter_vote -= ITEMS - 1 
            elec_votes.append(voter_vote)
        for voter in range(DISSIDENTS):
            if i == BAD_ELEC:
                elec_votes.append(fake_value)
            else:
                elec_votes.append(randint(0, ITEMS))

        votes.append(elec_votes)
    return votes, truth
                


           


    return []


if DISSIDENTS > VOTERS:
    print('More dissidents than voters!')
else:
    votes, truth = create_votes()
    get_vote(votes)
    for i in range(ELECTIONS):
        print("True value for election {} is {}".format(i+1, truth[i]))




