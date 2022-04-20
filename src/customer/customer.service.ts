import { Injectable } from '@nestjs/common';
import { CreateCustomerDto } from './dto/create-customer.dto';
import { UpdateCustomerDto } from './dto/update-customer.dto';
import { Customer, Prisma } from '@prisma/client'
import { PrismaService } from 'src/prisma.service';

@Injectable()
export class CustomerService {

  constructor(private prisma: PrismaService) { }

  create(data: Prisma.CustomerCreateInput): Promise<Customer> {
    return this.prisma.customer.create({
      data: data,
    });
  }

  findAll(params: {
    skip?: number,
    take?: number,
    cursor?: Prisma.CustomerWhereUniqueInput,
    where?: Prisma.CustomerWhereInput,
    orderBy?: Prisma.CustomerOrderByWithRelationInput,
  }): Promise<Customer[]> {
    return this.prisma.customer.findMany({
      ...params,
    });

  }

  findOne(customerWhereUniqueInput: Prisma.CustomerWhereUniqueInput): Promise<Customer | null> {
    return this.prisma.customer.findUnique({
      where: customerWhereUniqueInput,
    });

  }

  update(params: {
    where: Prisma.CustomerWhereUniqueInput,
    data: Prisma.CustomerUpdateInput,
  }): Promise<Customer> {
    const { where, data } = params;
    return this.prisma.customer.update({
      data,
      where,
    });
  }

  remove(where: Prisma.CustomerWhereUniqueInput): Promise<Customer> {
    return this.prisma.customer.delete({
      where,
    });
  }
}
