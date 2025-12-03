def adjacent_idx_generator(idx, length):
    return (idx + 1) % length

def opposite_idx_generator(idx, length):
    return (idx+length//2) % length

def match_sum(sequence, idx_generator):
    total = 0
    for i, digit in enumerate(sequence):
        if digit == sequence[idx_generator(i, len(sequence))]:
            total += int(digit)
    return total

def main():
    sequence = '91212129'
    total = match_sum(sequence, adjacent_idx_generator)
    print('Adjacent match sum:', total)

    total = match_sum(sequence, opposite_idx_generator)
    print('Opposite match sum:', total)

if __name__ == '__main__':
    main()