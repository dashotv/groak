#!/usr/bin/env bash

set -e

go run . scraper myanime

go run . downloader metube http://10.0.4.62:10003/add

# go run . page "swallowed star" "https://myanime.live/tag/swallowed-star/"
# go run . page "perfect world" "https://myanime.live/tag/perfect-world/"
# go run . page "a will eternal" "https://myanime.live/tag/a-will-eternal/"
# go run . page "battle through the heavens" "https://myanime.live/tag/battle-through-the-heavens/"
# go run . page "honor of kings chapter of glory" "https://myanime.live/tag/honor-of-kings-chapter-of-glory/"
# go run . page "against the gods" "https://myanime.live/tag/against-the-gods/"
# go run . page "glorious revenge of ye feng" "https://myanime.live/tag/dubu-wangu/"
# go run . page "the abyss game" "https://myanime.live/tag/shenyuan-youxi/"
# go run . page "shrouding the heavens" "https://myanime.live/tag/shrouding-the-heavens/"
# go run . page "a record of a mortals journey to immortality" "https://myanime.live/tag/fan-ren-xiu-xian-chuan/"
# go run . page "spare me great lord" "https://myanime.live/tag/spare-me-great-lord/"
# go run . page "law of devil" "https://myanime.live/tag/law-of-devil/"
# go run . page "the invincible" "https://myanime.live/tag/shi-fang-wu-sheng/"
# go run . page "the proud emperor of eternity" "https://myanime.live/tag/the-proud-emperor-of-eternity/"
# go run . page "legend of xianwu" "https://myanime.live/tag/legend-of-xianwu/"
# go run . page "martial god asura" "https://myanime.live/tag/martial-god-asura/"
# go run . page "100000 years of refining qi" "https://myanime.live/tag/100-000-years-of-refining-qi/page/10/"
go run . page "the eternal strife" "https://myanime.live/tag/bu-shi-bu-mie/"
# go run . page "the great ruler" "https://myanime.live/tag/the-great-ruler/"

go run . show
